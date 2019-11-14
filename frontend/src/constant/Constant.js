const PREFIX = ">";
const SUFFIX = "ago";

const TxType = {};
TxType.WITHDRAWADDRESS = 'WithdrawAddress';
TxType.BEGINUNBONDING = 'BeginUnbonding';
TxType.BEGINREDELEGATE = 'BeginRedelegate';
TxType.WITHDRAWDELEGATORREWARDSALL = 'WithdrawDelegatorRewardsAll';
TxType.WITHDRAWDELEGATORREWARD = 'WithdrawDelegatorReward';
TxType.TRANSFERS = 'Transfers';
TxType.STAKES = 'Stakes';
TxType.GOVERNANCE = 'Governance';
TxType.DECLARATIONS = 'Declarations';
TxType.DELEGATE = 'Delegate';
TxType.CREATEVALIDATOR = 'CreateValidator';
TxType.EDITVALIDATOR = 'EditValidator';
TxType.UNJAIL = 'Unjail';
TxType.SUBMITPROPOSAL = 'SubmitProposal';
TxType.DEPOSIT = 'Deposit';
TxType.VOTE = 'Vote';
TxType.SETWITHDRAWADDRESS = 'SetWithdrawAddress';

const ValidatorStatus = {};
ValidatorStatus.ACTIVE = 'active';
ValidatorStatus.JAILED = 'jailed';
ValidatorStatus.CANDIDATES = 'candidates';
ValidatorStatus.UNBONDED = 'Unbonded';
ValidatorStatus.UNBONDING = 'Unbonding';
ValidatorStatus.UNBONDED = 'Bonded';

const Denom = {};
Denom.IRISATTO = 'iris-atto';
Denom.IRIS = 'iris';
const ENVCONFIG = {};
ENVCONFIG.DEV = 'dev';
ENVCONFIG.QA = 'qa';
ENVCONFIG.STAGE = 'stage';
ENVCONFIG.TESTNET = 'testnet';
ENVCONFIG.MAINNET = 'mainnet';

const PARAMETER = {};
PARAMETER.EQUAL = 'eq';
PARAMETER.UNEQUAL = 'neq';

const CHAINNAME = 'iris';
const CHAINID = {};
CHAINID.IRISHUB = 'irishub';
CHAINID.FUXI = 'fuxi';
CHAINID.NYANCAT = 'nyancat';

const RADIXDENOM = {};

RADIXDENOM.IRISATTO = 'iris-atto';
RADIXDENOM.IRISATTONUMBER = '1000000000000000000';
RADIXDENOM.IRISMILLI = 'iris-milli';
RADIXDENOM.IRISMILLINUMBER = '1000000000000000';
RADIXDENOM.IRISMICRO = 'iris-micro';
RADIXDENOM.IRISMICRONUMBER = '1000000000000';
RADIXDENOM.IRISNANO = 'iris-nano';
RADIXDENOM.IRISNANONUMBER = '1000000000';
RADIXDENOM.IRISPICO = 'iris-pico';
RADIXDENOM.IRISPICONUMBER = '1000000';
RADIXDENOM.IRISFEMTO = 'iris-femto';
RADIXDENOM.IRISFEMTONUMBER = '1000';
RADIXDENOM.IRIS = 'iris';
RADIXDENOM.IRISNUMBER = '1';

export default {
  PREFIX,
  SUFFIX,
  TxType,
  ValidatorStatus,
  Denom,
  ENVCONFIG,
  CHAINNAME,
  PARAMETER,
  CHAINID,
  RADIXDENOM
};
