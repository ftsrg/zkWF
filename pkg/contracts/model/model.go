// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ModelMetaData contains all meta data concerning the Model contract.
var ModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initial_hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"initial_state\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"public_inputs\",\"type\":\"uint256[]\"}],\"name\":\"Verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentState\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"public_inputs\",\"type\":\"uint256[]\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516120d03803806120d083398101604081905261002e916100cb565b600481511461003b575f5ffd5b5f829055805161005290600190602084019061005a565b50505061019f565b828054828255905f5260205f20908101928215610093579160200282015b82811115610093578251825591602001919060010190610078565b5061009f9291506100a3565b5090565b5b8082111561009f575f81556001016100a4565b634e487b7160e01b5f52604160045260245ffd5b5f5f604083850312156100dc575f5ffd5b825160208401519092506001600160401b038111156100f9575f5ffd5b8301601f81018513610109575f5ffd5b80516001600160401b03811115610122576101226100b7565b604051600582901b90603f8201601f191681016001600160401b0381118282101715610150576101506100b7565b60405291825260208184018101929081018884111561016d575f5ffd5b6020850194505b8385101561019057845180825260209586019590935001610174565b50809450505050509250929050565b611f24806101ac5f395ff3fe60806040526004361061003e575f3560e01c806324fb01e814610042578063378aa701146100575780637cee919c146100815780637e4f7a8a1461009d575b5f5ffd5b610055610050366004611d3c565b6100cc565b005b348015610062575f5ffd5b5061006b61023d565b6040516100789190611e04565b60405180910390f35b34801561008c575f5ffd5b505f54604051908152602001610078565b3480156100a8575f5ffd5b506100bc6100b7366004611d3c565b610293565b6040519015158152602001610078565b5f5482825f8181106100e0576100e0611e46565b90506020020135146100f0575f5ffd5b3482826100fe600282611e6e565b81811061010d5761010d611e46565b905060200201351461011d575f5ffd5b8061012a60046007611e87565b610135906002611e87565b111561013f575f5ffd5b5f61014c85858585610293565b90508061015b5761015b611e9a565b8282600181811061016e5761016e611e46565b60200291909101355f555060075b61018860046007611e87565b8110156101da578383828181106101a1576101a1611e46565b9050602002013560016007836101b79190611e6e565b815481106101c7576101c7611e46565b5f9182526020909120015560010161017c565b505f83836101e9600182611e6e565b8181106101f8576101f8611e46565b9050602002013590505f81111561023557604051339082156108fc029083905f818181858888f19350505050158015610233573d5f5f3e3d5ffd5b505b505050505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561028957602002820191905f5260205f20905b815481526020019060010190808311610275575b5050505050905090565b5f60405161024081016102a584610598565b6102af85856105ab565b6102b8866105e7565b6102c1876105fd565b5f6102cd86868a6106d4565b90506102d8816109aa565b90506102e481896109fd565b90506102f08189610a60565b5060608201515f516020611eaf5f395f51905f525f516020611ecf5f395f51905f52610320846204000085611ce4565b086101c084015250610333818587610ab8565b6101a083015250610342610d18565b61034b86611a08565b6103548661195a565b61035d8661173d565b61036686611337565b61036f8661117e565b61037886610df7565b61020001519050611d34565b60405162461bcd60e51b815260206004820152601d60248201527f77726f6e67206e756d626572206f66207075626c696320696e707574730000006044820152606481fd5b60405162461bcd60e51b815260206004820152600c60248201526c06572726f72206d6f642065787609c1b6044820152606481fd5b60405162461bcd60e51b815260206004820152601260248201527132b93937b91032b19037b832b930ba34b7b760711b6044820152606481fd5b60405162461bcd60e51b815260206004820152601860248201527f696e707574732061726520626967676572207468616e207200000000000000006044820152606481fd5b60405162461bcd60e51b815260206004820152601060248201526f77726f6e672070726f6f662073697a6560801b6044820152606481fd5b60405162461bcd60e51b815260206004820152601660248201527537b832b734b733b9903134b3b3b2b9103a3430b7103960511b6044820152606481fd5b60405162461bcd60e51b815260206004820152600d60248201526c6572726f722070616972696e6760981b6044820152606481fd5b60405162461bcd60e51b815260206004820152600c60248201526b6572726f722076657269667960a01b6044820152606481fd5b60405162461bcd60e51b81526020600482015260146024820152736572726f722072616e646f6d2067656e206b7a6760601b6044820152606481fd5b600d81146105a8576105a8610384565b50565b5f5b818110156105e2575f516020611ecf5f395f51905f52833511156105d3576105d3610438565b602092909201916001016105ad565b505050565b6103008181146105f9576105f961047d565b5050565b61018081015f516020611ecf5f395f51905f5281351115610620576106206104b5565b506101a081015f516020611ecf5f395f51905f5281351115610644576106446104b5565b506101c081015f516020611ecf5f395f51905f5281351115610668576106686104b5565b506101e081015f516020611ecf5f395f51905f528135111561068c5761068c6104b5565b5061020081015f516020611ecf5f395f51905f52813511156106b0576106b06104b5565b5061026081015f516020611ecf5f395f51905f52813511156105f9576105f96104b5565b5f60405161024081016467616d6d6181527f0bc2470433fe75b1ea0f712df46413d1fff7e7bf4da0d4405222d56b235d117c60208201527f2cfe852c5d2b4c3a082f073563d82ba3b01ee59a46ac23df5acc67d1040e649f60408201527f0af0f405e09de97ed7b4739591616ed84ee532a899f079e5ad0da06916673a6560608201527f07585383535289d2ccbb4eb914e52a4897d404765024457e72734346325ba42060808201527f13a76652c4536eefcf83a7b1488c76fbeed15783a155e1e9e6778afcd44bcc0a60a08201527f09bac7c211128cd8f01caa00f0a576cb17a52070a18ad56d46cf3686508d000f60c08201527f105897e058e642a2f95d92ad09be2cc817ad74bf5ff266da7d0532e4d0eef98060e08201527f14190f62dd4d78e182c068ac838e12be373ae204f0a8659bf2e38e1e14e7674c6101008201527f234d6479f863e68ddbbf5b41cfa9af7ea597999fc0ac48f7390841d889eec2a96101208201527f0ed425ce26b9df0b06bb90c2410923be74448b7a110f2eb78c8d45f86b3b526a6101408201527f20cc1b4a0f58e05c83c7259b5a427864df7f2adaab042bb4470c871f589774b56101608201527f22fe42bf966f4e9f01139e1e42715fde0dde565122d98259da3ebf438238a3af6101808201527f02e6f2dbbfcb6bdb1f2d9f6e8a02b827d824bc42e67552473cf3f2ddc41dba4b6101a08201527f197992891a1d221948da54dc683491a86a0f878a6db64e1ae0609e05922376666101c08201527f06e990063a60c02428b9488de79b327fcf3f1f37e60b78c0aaa0488ab2ea0d846101e08201527f028d20ba485b4c172d0dd9057b5bba9122760eaf4dc6023ab5dca3acb9a78760610200820152610220810160208602808883379081019060c080878437506102c501905060208282601b820160025afa90508061098757610987610528565b5080519250505f516020611eaf5f395f51905f5282066040820152509392505050565b5f60405161024060405101636265746181528360208201526020816024601c840160025afa806109dc576109dc610528565b5080519250505f516020611eaf5f395f51905f528206602082015250919050565b5f60405161024060405101606564616c7068618252602082018681526020810190506040610220870182375060208282601b850160025afa905080610a4457610a44610528565b50515f516020611eaf5f395f51905f5281069091529392505050565b60405161024060405101637a657461815283602082015260c0808401604083013760208160e4601c840160025afa80610a9b57610a9b610528565b50515f516020611eaf5f395f51905f529006606091909101525050565b5f60405160608101516101c0820151915085610ad681878585610b2b565b5f92505f91505b85821015610b21575f516020611eaf5f395f51905f52853582510992505f516020611eaf5f395f51905f528385086020958601959094506001929092019101610add565b5050509392505050565b5f516020611eaf5f395f51905f527f30644259cd94e7dd5045d7a27013b7fcd21c9e3b7fa75222e7bda49b729b040183096001855f5b86811015610bcd575f516020611eaf5f395f51905f52835f516020611eaf5f395f51905f5203860882525f516020611eaf5f395f51905f527f19ddbcaf3a8d46c15c0176fbb5b95e4dc57088ff13f4d1bd84c6bfa57dcdc0e08409925060209190910190600101610b61565b50610bd9818789610c58565b5060019050855f5b86811015610c4e575f516020611eaf5f395f51905f52835f516020611eaf5f395f51905f52868551090982526020820191505f516020611eaf5f395f51905f527f19ddbcaf3a8d46c15c0176fbb5b95e4dc57088ff13f4d1bd84c6bfa57dcdc0e084099250600101610be1565b5050505050505050565b600183525f5f5b83811015610c9a5781850151828401515f516020611eaf5f395f51905f52818309905060208401935080848801525050600181019050610c5f565b506020810382019150808401935050610cc86020840160025f516020611eaf5f395f51905f52038551611ce4565b5f5b83811015610d115760208503945082515f516020611eaf5f395f51905f528651840984525f516020611eaf5f395f51905f52818409601f1990940193925050600101610cca565b5050505050565b604051610240604051016101c08201515f516020611eaf5f395f51905f5260015f516020611eaf5f395f51905f5203606085015108610d78837f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593efffffff83611ce4565b90505f516020611eaf5f395f51905f527f30644259cd94e7dd5045d7a27013b7fcd21c9e3b7fa75222e7bda49b729b0401820990505f516020611eaf5f395f51905f528282098451935091505f516020611eaf5f395f51905f52905082820990505f516020611eaf5f395f51905f528282099050806080840152505050565b60405161024081016101608201518152610180820151602082015261028083013560408201526102a08301356060820152610220830135608082015261024083013560a08201526102c083013560c08201526102e083013560e082015260608201516101008201526101e08201516101208201526020816101408360025afa80610e8357610e8361055c565b5f516020611eaf5f395f51905f5282510690508160408101925061028085013581526102a08501356020820152610ec083836102c0880184611c71565b6101608401610ed58484610220890184611c71565b6101408501610ee984610260890183611cb8565b6001855260026020860152805160408087019182529095908160608160075afa915081610f1857610f18610528565b60208101915081517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47038252610f5086828586611b68565b505083604085019450610f6d8560608801516102808a0184611bff565b5f516020611eaf5f395f51905f527f19ddbcaf3a8d46c15c0176fbb5b95e4dc57088ff13f4d1bd84c6bfa57dcdc0e060608801510995505f516020611eaf5f395f51905f528685099350610fc785856102c08a0184611c71565b610fd385828485611b68565b50602082810180517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd470381528251865291810151908501527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c260408501527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60608501527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60808501527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60a0850152905160c0840152805160e08401527f26186a2d65ee4d2f9c9a5b91f86597d35f192cd120caf7e935d8443d1938e23d6101008401527f30441fd1b5d3370482c42152a8899027716989a6996c2535bc9f7fee8aaef79e6101208401527f1970ea81dd6992adfbc571effb03503adbbb6a857f578403c6c40e22d65b3c026101408401527f054793348f12c0cf5622c340573cb277586319de359ab9389778f689786b1e4861016084015292506105e290508160405160205f6101808460085afa80611170576111706104f3565b505f51610200919091015250565b6040516102406040510160208101604082016101e084015180610160860160e087015181526101008701516101808801526101208701516101408801526111c986835f8b0184611c71565b6111dc826101808a016101408a01611cb8565b5f516020611eaf5f395f51905f5283830991506111fe868360408b0184611c71565b611211826101a08a016101408a01611cb8565b5f516020611eaf5f395f51905f528383099150611233868360808b0184611c71565b611246826101c08a016101408a01611cb8565b5f516020611eaf5f395f51905f5283830991507f0bc2470433fe75b1ea0f712df46413d1fff7e7bf4da0d4405222d56b235d117c86527f2cfe852c5d2b4c3a082f073563d82ba3b01ee59a46ac23df5acc67d1040e649f85526112ab84838884611c2a565b6112be826101e08a016101408a01611cb8565b5f516020611eaf5f395f51905f5283830991507f0af0f405e09de97ed7b4739591616ed84ee532a899f079e5ad0da06916673a6586527f07585383535289d2ccbb4eb914e52a4897d404765024457e72734346325ba420855261132384838884611c2a565b506102338161020089016101408901611cb8565b6040516467616d6d616102408201908152606082015161026083015260e08201516102808301526101008201516102a083015260c0836102c08401377f0bc2470433fe75b1ea0f712df46413d1fff7e7bf4da0d4405222d56b235d117c6101408201527f2cfe852c5d2b4c3a082f073563d82ba3b01ee59a46ac23df5acc67d1040e649f6101608201527f0af0f405e09de97ed7b4739591616ed84ee532a899f079e5ad0da06916673a65610180808301919091527f07585383535289d2ccbb4eb914e52a4897d404765024457e72734346325ba4206101a0808401919091526101208401516101c080850191909152918501356101e080850191909152908501356102008085019190915291850135610220840152808501356102408401529084013561026080840191909152840135610280830152601b906102859060209085018284860160025afa925050508061149357611493610528565b506101e00180515f516020611eaf5f395f51905f529006905250565b604051610240604051017f105897e058e642a2f95d92ad09be2cc817ad74bf5ff266da7d0532e4d0eef98081527f14190f62dd4d78e182c068ac838e12be373ae204f0a8659bf2e38e1e14e7674c6020820152611519604082016101808501358360e08601611bd4565b7f234d6479f863e68ddbbf5b41cfa9af7ea597999fc0ac48f7390841d889eec2a981527f0ed425ce26b9df0b06bb90c2410923be74448b7a110f2eb78c8d45f86b3b526a6020820152611579604082016101a08501358360e08601611c2a565b5f516020611eaf5f395f51905f526101a0840135610180850135097f20cc1b4a0f58e05c83c7259b5a427864df7f2adaab042bb4470c871f589774b582527f22fe42bf966f4e9f01139e1e42715fde0dde565122d98259da3ebf438238a3af60208301526115ef60408301828460e08701611c2a565b507f02e6f2dbbfcb6bdb1f2d9f6e8a02b827d824bc42e67552473cf3f2ddc41dba4b81527f197992891a1d221948da54dc683491a86a0f878a6db64e1ae0609e05922376666020820152611650604082016101c08501358360e08601611c2a565b7f06e990063a60c02428b9488de79b327fcf3f1f37e60b78c0aaa0488ab2ea0d8481527f028d20ba485b4c172d0dd9057b5bba9122760eaf4dc6023ab5dca3acb9a7876060208201526116ab604082018260e0850180611b68565b7f13a76652c4536eefcf83a7b1488c76fbeed15783a155e1e9e6778afcd44bcc0a81527f09bac7c211128cd8f01caa00f0a576cb17a52070a18ad56d46cf3686508d000f602082015261170660408201858360e08601611c2a565b6102208301358152610240830135602082015261172b60408201868360e08601611c2a565b610d118160a0840160e0850180611b68565b6040516020810151604082015160608301515f8401515f516020611eaf5f395f51905f5284610260880135095f516020611eaf5f395f51905f526101e088013586095f516020611eaf5f395f51905f52610180890135820890505f516020611eaf5f395f51905f5285820890505f516020611eaf5f395f51905f5261020089013587095f516020611eaf5f395f51905f526101a08a0135820890505f516020611eaf5f395f51905f5286820890505f516020611eaf5f395f51905f528284095f516020611eaf5f395f51905f5282820990505f516020611eaf5f395f51905f5285820990505f516020611eaf5f395f51905f52600580095f516020611eaf5f395f51905f52878a0998505f516020611eaf5f395f51905f526101808c01358a0894505f516020611eaf5f395f51905f5288860894505f516020611eaf5f395f51905f5260058a0993505f516020611eaf5f395f51905f526101a08c0135850893505f516020611eaf5f395f51905f5288850893505f516020611eaf5f395f51905f52818a099250505f516020611eaf5f395f51905f526101c08b0135830891505f516020611eaf5f395f51905f5287830891505f516020611eaf5f395f51905f5283850997505f516020611eaf5f395f51905f528289095f516020611eaf5f395f51905f52908103985085890997505f516020611eaf5f395f51905f5260808a01518908975061194e88828c6114af565b50505050505050505050565b604051600262040000016102406040510161197a81836060860151611ce4565b915061198f8183610140870160a08701611bff565b6119a281610100860160a0860180611b9e565b6119b1818360a0860180611bd4565b6119c38160c0860160a0860180611b9e565b6119da816101c085015160a0860160a08701611bd4565b505060c00180517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703905250565b6040515f5f516020611eaf5f395f51905f5260208301516101e08501350990505f516020611eaf5f395f51905f526040830151820890505f516020611eaf5f395f51905f52610180840135820890505f5f516020611eaf5f395f51905f5260208401516102008601350990505f516020611eaf5f395f51905f526040840151820890505f516020611eaf5f395f51905f526101a0850135820890505f5f516020611eaf5f395f51905f5260408501516101c08701350890505f516020611eaf5f395f51905f5282840992505f516020611eaf5f395f51905f528184099250505f516020611eaf5f395f51905f525f840151830991505f516020611eaf5f395f51905f52610260850135830991505f516020611eaf5f395f51905f526101a0840151830860808401519092505f516020611eaf5f395f51905f5290810391508183085f516020611eaf5f395f51905f52036101209390930192909252505050565b8151845260208201516020850152825160408501526020830151606085015260408160808660065afa80610d1157610d116103fe565b8151845260208201516020850152823560408501526020830135606085015260408160808660065afa80610d1157610d116103fe565b815184526020820151602085015282604085015260408160608660075afa80610d1157610d116103fe565b813584526020820135602085015282604085015260408160608660075afa80610d1157610d116103fe565b815184526020820151602085015282604085015260408460608660075afa815160408601526020820151606086015260408260808760065afa1680610d1157610d116103fe565b813584526020820135602085015282604085015260408460608660075afa815160408601526020820151606086015260408260808760065afa1680610d1157610d116103fe565b5f516020611eaf5f395f51905f52838335095f516020611eaf5f395f51905f5281835108825250505050565b602083526020808401526020604084015280606084015250806080830152505f516020611eaf5f395f51905f5260a08201525f60208260c08460055afa80611d2e57611d2e6103c9565b50505190565b949350505050565b5f5f5f5f60408587031215611d4f575f5ffd5b843567ffffffffffffffff811115611d65575f5ffd5b8501601f81018713611d75575f5ffd5b803567ffffffffffffffff811115611d8b575f5ffd5b876020828401011115611d9c575f5ffd5b60209182019550935085013567ffffffffffffffff811115611dbc575f5ffd5b8501601f81018713611dcc575f5ffd5b803567ffffffffffffffff811115611de2575f5ffd5b8760208260051b8401011115611df6575f5ffd5b949793965060200194505050565b602080825282518282018190525f918401906040840190835b81811015611e3b578351835260209384019390920191600101611e1d565b509095945050505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b81810381811115611e8157611e81611e5a565b92915050565b80820180821115611e8157611e81611e5a565b634e487b7160e01b5f52600160045260245ffdfe30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000130644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000a264697066735822122009213ff989bddff3c9f0bd7ea7d42fdab4eec8595a0afa1c38cf99f599ee6f0a64736f6c634300081c0033",
}

// ModelABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelMetaData.ABI instead.
var ModelABI = ModelMetaData.ABI

// ModelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModelMetaData.Bin instead.
var ModelBin = ModelMetaData.Bin

// DeployModel deploys a new Ethereum contract, binding an instance of Model to it.
func DeployModel(auth *bind.TransactOpts, backend bind.ContractBackend, initial_hash *big.Int, initial_state []*big.Int) (common.Address, *types.Transaction, *Model, error) {
	parsed, err := ModelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModelBin), backend, initial_hash, initial_state)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Model{ModelCaller: ModelCaller{contract: contract}, ModelTransactor: ModelTransactor{contract: contract}, ModelFilterer: ModelFilterer{contract: contract}}, nil
}

// Model is an auto generated Go binding around an Ethereum contract.
type Model struct {
	ModelCaller     // Read-only binding to the contract
	ModelTransactor // Write-only binding to the contract
	ModelFilterer   // Log filterer for contract events
}

// ModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelSession struct {
	Contract     *Model            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelCallerSession struct {
	Contract *ModelCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelTransactorSession struct {
	Contract     *ModelTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelRaw struct {
	Contract *Model // Generic contract binding to access the raw methods on
}

// ModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelCallerRaw struct {
	Contract *ModelCaller // Generic read-only contract binding to access the raw methods on
}

// ModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelTransactorRaw struct {
	Contract *ModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModel creates a new instance of Model, bound to a specific deployed contract.
func NewModel(address common.Address, backend bind.ContractBackend) (*Model, error) {
	contract, err := bindModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Model{ModelCaller: ModelCaller{contract: contract}, ModelTransactor: ModelTransactor{contract: contract}, ModelFilterer: ModelFilterer{contract: contract}}, nil
}

// NewModelCaller creates a new read-only instance of Model, bound to a specific deployed contract.
func NewModelCaller(address common.Address, caller bind.ContractCaller) (*ModelCaller, error) {
	contract, err := bindModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCaller{contract: contract}, nil
}

// NewModelTransactor creates a new write-only instance of Model, bound to a specific deployed contract.
func NewModelTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelTransactor, error) {
	contract, err := bindModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelTransactor{contract: contract}, nil
}

// NewModelFilterer creates a new log filterer instance of Model, bound to a specific deployed contract.
func NewModelFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelFilterer, error) {
	contract, err := bindModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelFilterer{contract: contract}, nil
}

// bindModel binds a generic wrapper to an already deployed contract.
func bindModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Model *ModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Model.Contract.ModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Model *ModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.Contract.ModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Model *ModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Model.Contract.ModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Model *ModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Model.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Model *ModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Model.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Model *ModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Model.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes proof, uint256[] public_inputs) view returns(bool success)
func (_Model *ModelCaller) Verify(opts *bind.CallOpts, proof []byte, public_inputs []*big.Int) (bool, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "Verify", proof, public_inputs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes proof, uint256[] public_inputs) view returns(bool success)
func (_Model *ModelSession) Verify(proof []byte, public_inputs []*big.Int) (bool, error) {
	return _Model.Contract.Verify(&_Model.CallOpts, proof, public_inputs)
}

// Verify is a free data retrieval call binding the contract method 0x7e4f7a8a.
//
// Solidity: function Verify(bytes proof, uint256[] public_inputs) view returns(bool success)
func (_Model *ModelCallerSession) Verify(proof []byte, public_inputs []*big.Int) (bool, error) {
	return _Model.Contract.Verify(&_Model.CallOpts, proof, public_inputs)
}

// GetCurrentHash is a free data retrieval call binding the contract method 0x7cee919c.
//
// Solidity: function getCurrentHash() view returns(uint256)
func (_Model *ModelCaller) GetCurrentHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getCurrentHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentHash is a free data retrieval call binding the contract method 0x7cee919c.
//
// Solidity: function getCurrentHash() view returns(uint256)
func (_Model *ModelSession) GetCurrentHash() (*big.Int, error) {
	return _Model.Contract.GetCurrentHash(&_Model.CallOpts)
}

// GetCurrentHash is a free data retrieval call binding the contract method 0x7cee919c.
//
// Solidity: function getCurrentHash() view returns(uint256)
func (_Model *ModelCallerSession) GetCurrentHash() (*big.Int, error) {
	return _Model.Contract.GetCurrentHash(&_Model.CallOpts)
}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns(uint256[])
func (_Model *ModelCaller) GetCurrentState(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Model.contract.Call(opts, &out, "getCurrentState")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns(uint256[])
func (_Model *ModelSession) GetCurrentState() ([]*big.Int, error) {
	return _Model.Contract.GetCurrentState(&_Model.CallOpts)
}

// GetCurrentState is a free data retrieval call binding the contract method 0x378aa701.
//
// Solidity: function getCurrentState() view returns(uint256[])
func (_Model *ModelCallerSession) GetCurrentState() ([]*big.Int, error) {
	return _Model.Contract.GetCurrentState(&_Model.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0x24fb01e8.
//
// Solidity: function update(bytes proof, uint256[] public_inputs) payable returns()
func (_Model *ModelTransactor) Update(opts *bind.TransactOpts, proof []byte, public_inputs []*big.Int) (*types.Transaction, error) {
	return _Model.contract.Transact(opts, "update", proof, public_inputs)
}

// Update is a paid mutator transaction binding the contract method 0x24fb01e8.
//
// Solidity: function update(bytes proof, uint256[] public_inputs) payable returns()
func (_Model *ModelSession) Update(proof []byte, public_inputs []*big.Int) (*types.Transaction, error) {
	return _Model.Contract.Update(&_Model.TransactOpts, proof, public_inputs)
}

// Update is a paid mutator transaction binding the contract method 0x24fb01e8.
//
// Solidity: function update(bytes proof, uint256[] public_inputs) payable returns()
func (_Model *ModelTransactorSession) Update(proof []byte, public_inputs []*big.Int) (*types.Transaction, error) {
	return _Model.Contract.Update(&_Model.TransactOpts, proof, public_inputs)
}
