package skywallet

import (
	"bytes"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/skycoin/skycoin/src/cipher"

	messages "github.com/fibercrypto/skywallet-protob/go"
)

// MessageCancel prepare Cancel request
func MessageCancel() ([][64]byte, error) {
	msg := &messages.Cancel{}
	data, err := proto.Marshal(msg)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_Cancel)
	return chunks, nil
}

// MessageButtonAck send this message (before user action) when the device expects the user to push a button
func MessageButtonAck() ([][64]byte, error) {
	buttonAck := &messages.ButtonAck{}
	data, err := proto.Marshal(buttonAck)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_ButtonAck)
	return chunks, nil
}

// MessagePassphraseAck send this message when the device expects receiving a Passphrase
func MessagePassphraseAck(passphrase string) ([][64]byte, error) {
	msg := &messages.PassphraseAck{
		Passphrase: proto.String(passphrase),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_PassphraseAck)
	return chunks, nil
}

// MessageWordAck send this message between each word of the seed (before user action) during device backup
func MessageWordAck(word string) ([][64]byte, error) {
	wordAck := &messages.WordAck{
		Word: proto.String(word),
	}
	data, err := proto.Marshal(wordAck)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_WordAck)
	return chunks, nil
}

// MessageCheckMessageSignature prepare CheckMessageSignature request
func MessageCheckMessageSignature(message, signature, address string) ([][64]byte, error) {
	msg := &messages.SkycoinCheckMessageSignature{
		Address:   proto.String(address),
		Message:   proto.String(message),
		Signature: proto.String(signature),
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_SkycoinCheckMessageSignature)
	return chunks, nil
}

// MessageAddressGen prepare MessageAddressGen request
func MessageAddressGen(addressN, startIndex uint32, confirmAddress bool) ([][64]byte, error) {
	skycoinAddress := &messages.SkycoinAddress{
		AddressN:       proto.Uint32(addressN),
		ConfirmAddress: proto.Bool(confirmAddress),
		StartIndex:     proto.Uint32(startIndex),
	}

	data, err := proto.Marshal(skycoinAddress)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_SkycoinAddress)
	return chunks, nil
}

// MessageAddressGenBip44 prepare MessageAddressGen request
func MessageAddressGenBip44(addressN, startIndex, coinType, account uint32, confirmAddress bool) ([][64]byte, error) {
	skycoinAddress := &messages.SkycoinAddress{
		ConfirmAddress: proto.Bool(confirmAddress),
		StartIndex:     proto.Uint32(0), // TODO remove this field from here
		AddressN:       proto.Uint32(0), // TODO remove this field from here
		Bip44Addr: &messages.Bip44AddrIndex{
			CoinType:          proto.Uint32(firstHardenedChild + coinType), // coinType'
			Account:           proto.Uint32(firstHardenedChild + account),  // account'
			Change:            proto.Uint32(0),
			AddressStartIndex: proto.Uint32(startIndex),
			AddressN:          proto.Uint32(addressN),
		},
	}

	data, err := proto.Marshal(skycoinAddress)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_SkycoinAddress)
	return chunks, nil
}

// MessageDeviceGetRawEntropy prepare GetEntropy request
func MessageDeviceGetRawEntropy(entropyBytes uint32) ([][64]byte, error) {
	getEntropy := &messages.GetRawEntropy{
		Size_: &entropyBytes,
	}

	data, err := proto.Marshal(getEntropy)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_GetRawEntropy)
	return chunks, nil
}

// MessageDeviceGetMixedEntropy prepare GetMixedEntropy request
func MessageDeviceGetMixedEntropy(entropyBytes uint32) ([][64]byte, error) {
	getEntropy := &messages.GetMixedEntropy{
		Size_: &entropyBytes,
	}

	data, err := proto.Marshal(getEntropy)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_GetMixedEntropy)
	return chunks, nil
}

// MessageApplySettings prepare MessageApplySettings request
func MessageApplySettings(usePassphrase *bool, label string, language string) ([][64]byte, error) {
	applySettings := &messages.ApplySettings{
		Label:    proto.String(label),
		Language: proto.String(language),
	}
	if usePassphrase != nil {
		applySettings.UsePassphrase = proto.Bool(*usePassphrase)
	}
	log.Println(applySettings)
	data, err := proto.Marshal(applySettings)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_ApplySettings)
	return chunks, nil
}

// MessageBackup prepare MessageBackup request
func MessageBackup() ([][64]byte, error) {
	backupDevice := &messages.BackupDevice{}
	data, err := proto.Marshal(backupDevice)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_BackupDevice)
	return chunks, nil
}

// MessageChangePin prepare MessageChangePin request
func MessageChangePin(remove *bool) ([][64]byte, error) {
	changePin := &messages.ChangePin{}
	if remove != nil {
		changePin.Remove = proto.Bool(*remove)
	}
	data, err := proto.Marshal(changePin)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_ChangePin)
	return chunks, nil
}

// MessageConnected prepare MessageConnected request
func MessageConnected() ([][64]byte, error) {
	msgRaw := &messages.Ping{}
	data, err := proto.Marshal(msgRaw)
	if err != nil {
		return [][64]byte{}, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_Ping)
	return chunks, nil
}

// MessageFirmwareErase prepare MessageFirmwareErase request
func MessageFirmwareErase(payload []byte) ([][64]byte, error) {
	deviceFirmwareErase := &messages.FirmwareErase{
		Length: proto.Uint32(uint32(len(payload))),
	}

	erasedata, err := proto.Marshal(deviceFirmwareErase)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(erasedata, messages.MessageType_MessageType_FirmwareErase)
	return chunks, nil
}

// MessageFirmwareUpload prepare MessageFirmwareUpload request
func MessageFirmwareUpload(payload []byte, hash [32]byte) ([][64]byte, error) {
	deviceFirmwareUpload := &messages.FirmwareUpload{
		Payload: payload,
		Hash:    hash[:],
	}

	uploaddata, err := proto.Marshal(deviceFirmwareUpload)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(uploaddata, messages.MessageType_MessageType_FirmwareUpload)
	return chunks, nil
}

// MessageGetFeatures prepare MessageGetFeatures request
func MessageGetFeatures() ([][64]byte, error) {
	featureMsg := &messages.GetFeatures{}
	data, err := proto.Marshal(featureMsg)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_GetFeatures)
	return chunks, nil
}

// MessageGenerateMnemonic prepare MessageGenerateMnemonic request
func MessageGenerateMnemonic(wordCount uint32, usePassphrase bool) ([][64]byte, error) {
	skycoinGenerateMnemonic := &messages.GenerateMnemonic{
		PassphraseProtection: proto.Bool(usePassphrase),
		WordCount:            proto.Uint32(wordCount),
	}

	data, err := proto.Marshal(skycoinGenerateMnemonic)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_GenerateMnemonic)
	return chunks, nil
}

// MessageRecovery prepare MessageRecovery request
func MessageRecovery(wordCount uint32, usePassphrase *bool, dryRun bool) ([][64]byte, error) {
	recoveryDevice := &messages.RecoveryDevice{
		WordCount: proto.Uint32(wordCount),
		DryRun:    proto.Bool(dryRun),
	}
	if usePassphrase != nil {
		recoveryDevice.PassphraseProtection = proto.Bool(*usePassphrase)
	}
	data, err := proto.Marshal(recoveryDevice)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_RecoveryDevice)

	return chunks, nil
}

// MessageSetMnemonic prepare MessageSetMnemonic request
func MessageSetMnemonic(mnemonic string) ([][64]byte, error) {
	skycoinSetMnemonic := &messages.SetMnemonic{
		Mnemonic: proto.String(mnemonic),
	}

	data, err := proto.Marshal(skycoinSetMnemonic)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_SetMnemonic)
	return chunks, nil
}

// MessageSignMessage prepare MessageSignMessage request
func MessageSignMessage(addressIndex int, message string) ([][64]byte, error) {
	skycoinSignMessage := &messages.SkycoinSignMessage{
		AddressN: proto.Uint32(uint32(addressIndex)),
		Message:  proto.String(message),
	}

	data, err := proto.Marshal(skycoinSignMessage)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_SkycoinSignMessage)
	return chunks, nil
}

// MessageSignMessageBip44 prepare MessageSignMessage request
func MessageSignMessageBip44(startIndex, addressN, coinType, account uint32, message string) ([][64]byte, error) {
	skycoinSignMessage := &messages.SkycoinSignMessage{
		AddressN: proto.Uint32(uint32(0)), // TODO remove this field
		Message:  proto.String(message),
		Bip44Addr: &messages.Bip44AddrIndex{
			CoinType:          proto.Uint32(firstHardenedChild + coinType), // coinType'
			Account:           proto.Uint32(firstHardenedChild + account),  // account'
			Change:            proto.Uint32(0),
			AddressStartIndex: proto.Uint32(startIndex),
			AddressN:          proto.Uint32(addressN),
		},
	}

	data, err := proto.Marshal(skycoinSignMessage)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_SkycoinSignMessage)
	return chunks, nil
}

// MessageTransactionSign prepare MessageTransactionSign request
func MessageTransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput) ([][64]byte, error) {
	skycoinTransactionSignMessage := &messages.TransactionSign{
		NbIn:           proto.Uint32(uint32(len(inputs))),
		NbOut:          proto.Uint32(uint32(len(outputs))),
		TransactionIn:  inputs,
		TransactionOut: outputs,
	}
	log.Println(skycoinTransactionSignMessage)

	data, err := proto.Marshal(skycoinTransactionSignMessage)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_TransactionSign)
	return chunks, nil
}

// MessageTransactionSignBip44 prepare MessageTransactionSign request
func MessageTransactionSignBip44(coinType, account uint32, inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput) ([][64]byte, error) {
	for idxInput := range inputs {
		inputs[idxInput].Bip44Addr = &messages.Bip44AddrIndex{
			CoinType:          proto.Uint32(firstHardenedChild + coinType), // coinType'
			Account:           proto.Uint32(firstHardenedChild + account),  // account'
			Change:            proto.Uint32(0),
			AddressStartIndex: proto.Uint32(*inputs[idxInput].Index),
			AddressN:          proto.Uint32(1),
		}
	}
	for idxOutput := range outputs {
		outputs[idxOutput].Bip44Addr = &messages.Bip44AddrIndex{
			CoinType:          proto.Uint32(firstHardenedChild + coinType), // coinType'
			Account:           proto.Uint32(firstHardenedChild + account),  // account'
			Change:            proto.Uint32(0),
			AddressStartIndex: proto.Uint32(*outputs[idxOutput].AddressIndex),
			AddressN:          proto.Uint32(1),
		}
	}
	skycoinTransactionSignMessage := &messages.TransactionSign{
		NbIn:           proto.Uint32(uint32(len(inputs))),
		NbOut:          proto.Uint32(uint32(len(outputs))),
		TransactionIn:  inputs,
		TransactionOut: outputs,
	}

	data, err := proto.Marshal(skycoinTransactionSignMessage)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_TransactionSign)
	return chunks, nil
}

// MessageWipe prepare MessageWipe request
func MessageWipe() ([][64]byte, error) {
	wipeDevice := &messages.WipeDevice{}
	data, err := proto.Marshal(wipeDevice)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_WipeDevice)
	return chunks, nil
}

// MessagePinMatrixAck prepare MessagePinMatrixAck request
func MessagePinMatrixAck(p string) ([][64]byte, error) {
	pinAck := &messages.PinMatrixAck{
		Pin: proto.String(p),
	}
	data, err := proto.Marshal(pinAck)
	if err != nil {
		return [][64]byte{}, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_PinMatrixAck)
	return chunks, nil
}

// MessageEntropyAck prepare MessageEntropyAck request
func MessageEntropyAck(bufferSize int) ([][64]byte, error) {
	buffer := cipher.RandByte(bufferSize)
	if len(buffer) != bufferSize {
		return nil, fmt.Errorf("required %d bytes but got %d", bufferSize, len(buffer))
	}
	entropyAck := &messages.EntropyAck{
		Entropy: buffer,
	}
	data, err := proto.Marshal(entropyAck)
	if err != nil {
		return nil, err
	}
	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_EntropyAck)
	return chunks, nil
}

// MessageInitialize prepare MessageInitialize request
func MessageInitialize() ([][64]byte, error) {
	initialize := &messages.Initialize{}
	data, err := proto.Marshal(initialize)
	if err != nil {
		return nil, err
	}

	chunks := makeSkyWalletMessage(data, messages.MessageType_MessageType_Initialize)
	return chunks, nil
}

// MessageSimulateButtonPress prespares a emulator button press simulation button
func MessageSimulateButtonPress(buttonType ButtonType) (*bytes.Buffer, error) {
	switch buttonType {
	case ButtonLeft, ButtonRight, ButtonBoth:
		msg := new(bytes.Buffer)
		msg.Write([]byte{0, 1, 2, 3, 4, byte(buttonType)})

		return msg, nil
	default:
		return nil, fmt.Errorf("invalid button type: %d", buttonType)
	}

}
