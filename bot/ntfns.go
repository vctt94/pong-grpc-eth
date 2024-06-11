package bot

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/companyzero/bisonrelay/clientrpc/types"
	"github.com/vctt94/pong-bisonrelay/pongrpc/grpc/pong"
)

func (b *Bot) kxNtfns(ctx context.Context) error {
	var ksr types.KXStreamRequest
	var ackReq types.AckRequest
	var ackRes types.AckResponse
	for {
		stream, err := b.chatService.KXStream(ctx, &ksr)
		if errors.Is(err, context.Canceled) {
			// Program is done.
			return err
		}
		if err != nil {
			b.kxLog.Warnf("Error while obtaining KX stream: %v", err)
			time.Sleep(time.Second) // Wait to try again.
			continue
		}
		b.kxLog.Info("Listening for kxs...")
		for {
			var pm types.KXCompleted
			err := stream.Recv(&pm)
			if errors.Is(err, context.Canceled) {
				// Program is done.
				return err
			}
			if err != nil {
				b.kxLog.Warnf("Error while receiving stream: %v", err)
				break
			}
			ksr.UnackedFrom = pm.SequenceId
			ackReq.SequenceId = pm.SequenceId
			if err = b.chatService.AckKXCompleted(ctx, &ackReq, &ackRes); err != nil {
				b.kxLog.Errorf("failed to acknowledge kx: %v", err)
				break
			}
			b.kxChan <- pm
		}
	}
}

func (b *Bot) pmNtfns(ctx context.Context) error {
	var ackRes types.AckResponse
	var ackReq types.AckRequest
	for {
		// Keep requesting a new stream if the connection breaks.
		streamReq := types.PMStreamRequest{UnackedFrom: ackReq.SequenceId}
		stream, err := b.chatService.PMStream(ctx, &streamReq)
		if errors.Is(err, context.Canceled) {
			// Program is done.
			return err
		}
		if err != nil {
			b.pmLog.Errorf("failed to get PM stream: %v", err)
			time.Sleep(time.Second) // Wait to try again.
			continue
		}

		b.pmLog.Info("Listening for private messages...")
		for {
			var pm types.ReceivedPM
			err := stream.Recv(&pm)
			clientID := hex.EncodeToString(pm.GetUid())

			if errors.Is(err, context.Canceled) {
				// Program is done.
				return err
			}
			if err != nil {
				b.pmLog.Errorf("failed to receive from stream: %v", err)
				break
			}
			ackReq.SequenceId = pm.SequenceId
			err = b.chatService.AckReceivedPM(ctx, &ackReq, &ackRes)
			if err != nil {
				b.pmLog.Errorf("failed to acknowledge received gc: %v", err)
				break
			}
			if pm.Msg.Message == "!init" {
				// b.handleInit(clientID)
				return nil
			}
			// Check if the message is the !ready command
			if pm.Msg.Message == "!ready" {
				b.handleReadySignal(clientID)
				return nil
			}
			b.pmChan <- pm
		}
	}
}

// func (b *Bot) handleInit(clientId string) {
// 	// Call the SignalReady method on the server
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := b.gameServer.Init(ctx, &pong.GameStartedStreamRequest{
// 		ClientId: clientId,
// 	})

// 	if err != nil {
// 		b.pmLog.Errorf("failed to signal ready for user %s: %v", clientId, err)
// 	}
// }

func (b *Bot) handleReadySignal(clientId string) {
	// Call the SignalReady method on the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := b.gameServer.SignalReady(ctx, &pong.SignalReadyRequest{
		ClientId: clientId,
	})
	if err != nil {
		b.pmLog.Errorf("failed to signal ready for user %s: %v", clientId, err)
	}
}

func (b *Bot) tipProgress(ctx context.Context) error {
	var tpr types.TipProgressRequest
	var ackReq types.AckRequest
	var ackRes types.AckResponse
	for {
		stream, err := b.paymentService.TipProgress(ctx, &tpr)
		if errors.Is(err, context.Canceled) {
			// Program is done.
			return err
		}
		if err != nil {
			b.tipLog.Warnf("Error while creating tip progress stream: %v", err)
			time.Sleep(time.Second) // Wait to try again.
			continue
		}
		b.tipLog.Info("Listening for tip progress...")
		for {
			var pm types.TipProgressEvent
			err := stream.Recv(&pm)
			if errors.Is(err, context.Canceled) {
				// Program is done.
				return err
			}
			if err != nil {
				b.tipLog.Warnf("Error while receiving stream: %v", err)
				break
			}
			tpr.UnackedFrom = pm.SequenceId
			ackReq.SequenceId = pm.SequenceId
			if err = b.paymentService.AckTipProgress(ctx, &ackReq, &ackRes); err != nil {
				b.tipLog.Errorf("Failed to acknowledge tip progress: %v", err)
				break
			}
			b.tipChan <- pm
		}
	}
}