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

export default {
  PREFIX,
  SUFFIX,
  TxType,
  ValidatorStatus,
  Denom,
};
