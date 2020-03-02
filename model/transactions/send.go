package transactions

import (
	"github.com/isarq/nem-sdk-go/com/requests"
	"github.com/isarq/nem-sdk-go/extras"
	"github.com/isarq/nem-sdk-go/model"
	"github.com/isarq/nem-sdk-go/utils"
	"github.com/pkg/errors"
)

type Common struct {
	Password, PrivateKey string
	IsHW                 bool
}

// Serialize a transaction and broadcast it to the network
// param common - A common struct
// param entity - A prepared transaction struct
// param endpoint - An NIS endpoint struct
// return - An announce transaction promise of the com.requests service
func Send(common Common, entity interface{}, endpoint *requests.Client) (*requests.NemAnnounceResult, error) {
	if extras.IsEmpty(common) || extras.IsEmpty(entity) || extras.IsEmpty(endpoint) {
		return nil, errors.New("Missing parameter !")
	}
	if len(common.PrivateKey) != 64 && len(common.PrivateKey) != 66 {
		return nil, errors.New("Invalid private key, length must be 64 or 66 characters !")
	}
	if !utils.IsHexadecimal(common.PrivateKey) {
		return nil, errors.New("Private key must be hexadecimal only !")
	}
	kp, err := model.KeyPairCreate(common.PrivateKey)
	if err != nil {
		return nil, err
	}

	result := utils.SerializeTransaction(entity)
	signature, err := kp.Sign(result)
	if err != nil {
		return nil, err
	}

	obj := requests.RequestAnnounce{
		Data:      utils.Bt2Hex([]byte(result)),
		Signature: utils.Bt2Hex(signature),
	}
	return endpoint.Announce(obj)
}
