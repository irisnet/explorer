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
const CoinsMap = {};


CoinsMap['iris'] = 1;
CoinsMap['iris-milli'] = 1000;
CoinsMap['iris-micro'] = 1000000;
CoinsMap['iris-nano'] = 1000000000;
CoinsMap['iris-pico'] = 1000000000000;
CoinsMap['iris-femto'] = 1000000000000000;
CoinsMap['iris-atto'] = 1000000000000000000;

export default {
  PREFIX,
  SUFFIX,
  TxType,
  ValidatorStatus,
  Denom,
  CoinsMap
};
