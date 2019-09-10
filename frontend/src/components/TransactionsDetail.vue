<template>
  <div class="transactions_detail_wrap">
    <div :class="transactionsDetailWrap">
      <p
        class="transaction_information_content_title"
        style="height: 0.7rem; line-height: 0.7rem;"
      >Transaction Information</p>
      <div class="transactions_detail_information_wrap" ref="valueInformation">
        <div class="information_props_wrap">
          <span class="information_props">TxHash :</span>
          <div style="display: flex;">
            <span class="information_value">{{info.TxHash}}</span>
            <m-clip v-if="info.TxHash" style="margin-left: 8px;" :text="info.TxHash"></m-clip>
          </div>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Block :</span>
          <span class="information_value link_active_style">
            <router-link :to="`/block/${info.Block}`">{{info.Block}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Status :</span>
          <span class="information_value information_value_fixed">
            <span :class="{'fail_status': info.Status === 'Fail' }">{{info.Status}}</span>
            <div
              class="info_icon_div question_icon_div"
              v-if="info.Status === 'Fail' && failInfo"
              v-table-tooltip="{show: true, container: $refs.valueInformation}"
            >
              <div class="tooltip_span">
                <div>{{failInfo}}</div>
                <i></i>
              </div>
            </div>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Timestamp :</span>
          <span class="information_value" v-show="ageValue">{{ageValue}} ({{info.Timestamp}})</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Fee :</span>
          <span class="information_value">{{info.Fee || '--'}}</span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Gas Used :</span>
          <span class="information_value information_value_fixed">
            <span>{{info.GasUsed || '--'}}</span>
            <div
              class="info_icon_div"
              v-if="gasPrice || info.GasUsed || gasWanted || gasLimit"
              v-table-tooltip="{show: true, container: $refs.valueInformation}"
            >
              <div class="tooltip_span">
                <div>
                  <p>Gas Price : {{gasPrice ? gasPrice + ' Nano' : '--'}}</p>
                  <p>Gas Used : {{info.GasUsed || '--'}}</p>
                  <p>Gas Wanted : {{gasWanted || '--'}}</p>
                  <p>Gas Limit : {{gasLimit || '--'}}</p>
                </div>
                <i></i>
              </div>
            </div>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Signer :</span>
          <span class="information_value link_active_style">
            <router-link :to="addressRoute(info.Signer)">{{info.Signer}}</router-link>
          </span>
        </div>
        <div class="information_props_wrap">
          <span class="information_props">Memo :</span>
          <span class="information_value">
            <pre class="information_pre">{{info.Memo || '--'}}</pre>
          </span>
        </div>
      </div>
      <p
        class="transaction_information_content_title"
        style="height: 0.7rem; line-height: 0.7rem;"
        v-show="txTypeSignData.length > 0"
      >Transaction Message</p>
      <div
        class="transactions_detail_information_wrap"
        v-for="(item, i) in txTypeSignData"
        :key="i"
      >
        <!-- <div class="information_props_wrap">
                    <span class="information_props">TxType :</span>
                    <span class="information_value">{{typeValue}}</span>
        </div>-->
        <div class="information_props_wrap" v-for="(v, k) in item" :key="k" v-show="!v.hide">
          <span class="information_props">{{v.real_k || k}} :</span>
          <template v-if="v.v !== '' && v.v !== undefined">
            <template v-if="addrFields.includes(k)">
              <template v-if="Array.isArray(v.v)">
                <div v-show="v.v.length">
                  <p v-for="it in v.v" :key="it" class="information_value link_active_style">
                    <router-link :to="addressRoute(it)">{{it}}</router-link>
                  </p>
                </div>
                <p class="information_value" v-show="!v.v.length">--</p>
              </template>
              <template v-else>
                <span class="information_value link_active_style">
                  <router-link :to="addressRoute(v.v)">{{v.f ? (item[v.f].v || v.v) : v.v}}</router-link>
                </span>
              </template>
            </template>
            <template v-else-if="k === 'Proposal ID'">
              <span class="information_value link_active_style">
                <router-link :to="`/ProposalsDetail/${v.v}`">{{v.v}}</router-link>
              </span>
            </template>
            <template v-else-if="k === 'Identity'">
              <span v-show="v.v && v.identityUrl" class="information_value link_active_style">
                <a :href="v.identityUrl" target="_blank">{{v.v}}</a>
              </span>
              <span v-show="v.v && !v.identityUrl" class="information_value">{{v.v}}</span>
            </template>
            <template v-else-if="k === 'Website'">
              <span
                v-show="v.v !== '[do-not-modify]'"
                @click="openUrl(v.v)"
                class="information_value link_active_style"
              >{{v.v}}</span>
              <span v-show="v.v === '[do-not-modify]'" class="information_value">{{v.v}}</span>
            </template>
            <template v-else-if="k === 'Software'">
              <span
                v-show="v.v && v.v.startsWith('http')"
                class="information_value link_active_style"
              >
                <a :href="v.v" target="_blank">{{v.v}}</a>
              </span>
              <span v-show="v.v && !v.v.startsWith('http')" class="information_value">{{v.v}}</span>
            </template>
            <template v-else-if="k === 'Commission Rate'">
              <span class="information_value information_value_fixed">
                <span>{{v.v || '--'}}</span>
                <div
                  class="info_icon_div"
                  v-if="typeValue === 'CreateValidator' && (item['Max Rate'].v || item['Max Change Rate'].v)"
                  v-table-tooltip="{show: true, container: $refs.valueInformation}"
                >
                  <div class="tooltip_span">
                    <div>
                      <p>Max Rate : {{item['Max Rate'].v || '--'}}</p>
                      <p>Max Change Rate : {{item['Max Change Rate'].v || '--'}}</p>
                    </div>
                    <i></i>
                  </div>
                </div>
              </span>
            </template>
            <template v-else>
              <span class="information_value">{{v.v}}</span>
            </template>
          </template>
          <template v-else>
            <span class="information_value">--</span>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Tools from "../util/Tools";
import Service from "../service";
import Constant from "../constant/Constant";
import MClip from "./commonComponents/MClip";
import axios from "../util/axios";
import Constants from "../constant/Constant";
export default {
  data() {
    return {
      transactionsDetailWrap: "personal_computer_transactions_detail",
      typeValue: "",
      gasLimit: "",
      gasPrice: "",
      gasWanted: "",
      failInfo: "",
      ageValue: "",
      transactionDetailTimer: null,
      addrFields: [
        "From",
        "To",
        "Operator Address",
        "Owner Address",
        "Original Address",
        "New Address",
        "Proposer",
        "Depositor",
        "Voter",
        "Owner",
        "Original Owner",
        "New Owner"
      ],
      info: {
        TxHash: "",
        Block: "",
        Status: "",
        Timestamp: "",
        Fee: "",
        GasUsed: "",
        Signer: "",
        Memo: ""
      },
      txTypeArr: {
        type_0: ["WithdrawDelegatorRewardsAll", "WithdrawValidatorRewardsAll"],
        type_1: ["Transfer", "Delegate", "WithdrawDelegatorReward"],
        type_2: ["Burn"],
        type_3: ["CreateValidator"],
        type_17: ["EditValidator"],
        type_4: ["SetWithdrawAddress"],
        type_5: ["Unjail"],
        type_6: ["SubmitProposal"],
        type_7: ["Deposit"],
        type_8: ["Vote"],
        type_9: ["IssueToken"],
        type_10: ["EditToken"],
        type_11: ["MintToken"],
        type_12: ["TransferTokenOwner"],
        type_13: ["CreateGateway", "EditGateway"],
        type_14: ["TransferGatewayOwner"],
        type_15: ["RequestRand"],
        type_16: ["SetMemoRegexp"],
        type_18: ["BeginUnbonding"],
        type_19: ["BeginRedelegate"]
      },
      txTypeSign: "",
      type_0: {
        TxType: {
          k: "type",
          v: ""
        },
        From: {
          k: "from",
          v: [],
          f: "from_moniker"
        },
        Amount: {
          k: "amount",
          v: ""
        },
        To: {
          k: "to",
          v: "",
          f: "to_moniker"
        },
        from_moniker: {
          hide: true,
          k: "from_moniker",
          v: ""
        },
        to_moniker: {
          hide: true,
          k: "to_moniker",
          v: ""
        }
      },
      type_1: {
        TxType: {
          k: "type",
          v: ""
        },
        From: {
          k: "from",
          v: "",
          f: "from_moniker"
        },
        Amount: {
          k: "amount",
          v: ""
        },
        To: {
          k: "to",
          v: "",
          f: "to_moniker"
        },
        from_moniker: {
          hide: true,
          k: "from_moniker",
          v: ""
        },
        to_moniker: {
          hide: true,
          k: "to_moniker",
          v: ""
        }
      },
      type_2: {
        TxType: {
          k: "type",
          v: ""
        },
        From: {
          k: "from",
          v: ""
        },
        Amount: {
          k: "amount",
          v: ""
        }
      },
      type_3: {
        TxType: {
          k: "type",
          v: ""
        },
        "Operator Address": {
          k: "operator_addr",
          v: ""
        },
        Moniker: {
          k: "moniker",
          v: ""
        },
        Identity: {
          k: "identity",
          v: "",
          identityUrl: ""
        },
        "Self-Bonded": {
          k: "",
          v: ""
        },
        "Owner Address": {
          k: "owner",
          v: ""
        },
        "Consensus Pubkey": {
          k: "pubkey",
          v: ""
        },
        "Commission Rate": {
          k: "rate",
          v: ""
        },
        "Max Rate": {
          hide: true,
          k: "max_rate",
          v: ""
        },
        "Max Change Rate": {
          hide: true,
          k: "max_change_rate",
          v: ""
        },
        Website: {
          k: "website",
          v: ""
        },
        Details: {
          k: "details",
          v: ""
        }
      },
      type_17: {
        TxType: {
          k: "type",
          v: ""
        },
        "Operator Address": {
          k: "operator_addr",
          v: ""
        },
        Moniker: {
          k: "moniker",
          v: ""
        },
        Identity: {
          k: "identity",
          v: "",
          identityUrl: ""
        },
        "Commission Rate": {
          k: "rate",
          v: ""
        },
        Website: {
          k: "website",
          v: ""
        },
        Details: {
          k: "details",
          v: ""
        }
      },
      type_4: {
        TxType: {
          k: "type",
          v: ""
        },
        "Original Address": {
          k: "from",
          v: ""
        },
        "New Address": {
          k: "to",
          v: ""
        }
      },
      type_5: {
        TxType: {
          k: "type",
          v: ""
        },
        "Operator Address": {
          k: "operator_addr",
          v: ""
        }
      },
      type_6: {
        TxType: {
          k: "type",
          v: ""
        },
        Proposer: {
          k: "from",
          v: ""
        },
        Title: {
          k: "title",
          v: ""
        },
        "Initial Deposit": {
          k: "Initial Deposit",
          v: ""
        },
        Description: {
          k: "description",
          v: ""
        },
        "Proposal Type": {
          k: "proposal_type",
          v: ""
        }
        // "Deposit Endtime": {
        //     k: "Deposit Endtime",
        //     v: ""
        // }
        // "Total Deposit": {
        //     k: "Total Deposit",
        //     v: ""
        // },
        // "Voting Starttime": {
        //     k: "Voting Starttime",
        //     v: ""
        // },
        // "Voting Endtime": {
        //     k: "Voting Endtime",
        //     v: ""
        // }
      },
      Parameter: {
        Parameter: {
          k: "param",
          v: ""
        }
      },
      SoftwareUpgrade: {
        Software: {
          k: "software",
          v: ""
        },
        Version: {
          k: "version",
          v: ""
        },
        "Switch Height": {
          k: "switch_height",
          v: ""
        },
        Treshold: {
          k: "treshold",
          v: ""
        }
      },
      type_7: {
        TxType: {
          k: "type",
          v: ""
        },
        Depositor: {
          k: "from",
          v: ""
        },
        "Proposal ID": {
          k: "proposal_id",
          v: ""
        },
        Deposit: {
          k: "",
          v: ""
        }
      },
      type_8: {
        TxType: {
          k: "type",
          v: ""
        },
        Voter: {
          k: "from",
          v: ""
        },
        "Proposal ID": {
          k: "proposal_id",
          v: ""
        },
        Option: {
          k: "option",
          v: ""
        }
      },
      type_9: {
        TxType: {
          k: "type",
          v: ""
        },
        Family: {
          k: "family",
          v: ""
        },
        Source: {
          k: "source",
          v: ""
        },
        Gateway: {
          k: "gateway",
          v: ""
        },
        Symbol: {
          k: "symbol",
          v: ""
        },
        "Canonical Symbol": {
          k: "canonical_symbol",
          v: ""
        },
        Name: {
          k: "name",
          v: ""
        },
        Decimal: {
          k: "decimal",
          v: ""
        },
        SymbolMinAlias: {
          k: "min_unit_alias",
          v: ""
        },
        InitialSupply: {
          k: "initial_supply",
          v: ""
        },
        MaxSupply: {
          k: "max_supply",
          v: ""
        },
        Mintable: {
          k: "mintable",
          v: ""
        },
        Owner: {
          k: "owner",
          v: ""
        }
      },
      type_10: {
        TxType: {
          k: "type",
          v: ""
        },
        TokenId: {
          k: "token_id",
          v: ""
        },
        "Canonical Symbol": {
          k: "canonical_symbol",
          v: ""
        },
        Name: {
          k: "name",
          v: ""
        },
        SymbolMinAlias: {
          k: "min_unit_alias",
          v: ""
        },
        MaxSupply: {
          k: "max_supply",
          v: ""
        },
        Mintable: {
          k: "mintable",
          v: ""
        },
        Owner: {
          k: "owner",
          v: ""
        }
      },
      type_11: {
        TxType: {
          k: "type",
          v: ""
        },
        TokenId: {
          k: "token_id",
          v: ""
        },
        Owner: {
          k: "owner",
          v: ""
        },
        Amount: {
          k: "amount",
          v: ""
        },
        To: {
          k: "to",
          v: ""
        }
      },
      type_12: {
        TxType: {
          k: "type",
          v: ""
        },
        TokenId: {
          k: "token_id",
          v: ""
        },
        "Original Owner": {
          k: "src_owner",
          v: ""
        },
        "New Owner": {
          k: "dst_owner",
          v: ""
        }
      },
      type_13: {
        TxType: {
          k: "type",
          v: ""
        },
        Owner: {
          k: "owner",
          v: ""
        },
        Moniker: {
          k: "moniker",
          v: ""
        },
        Identity: {
          k: "identity",
          v: "",
          identityUrl: ""
        },
        Details: {
          k: "details",
          v: ""
        },
        Website: {
          k: "website",
          v: ""
        }
      },
      type_14: {
        TxType: {
          k: "type",
          v: ""
        },
        "Original Owner": {
          k: "owner",
          v: ""
        },
        Moniker: {
          k: "moniker",
          v: ""
        },
        "New Owner": {
          k: "to",
          v: ""
        }
      },
      type_15: {
        TxType: {
          k: "type",
          v: ""
        },
        "Block Interval": {
          k: "block-interval",
          v: ""
        },
        "Request ID": {
          k: "request-id",
          v: ""
        },
        "Rand Height": {
          k: "rand-height",
          v: ""
        }
      },
      type_16: {
        TxType: {
          k: "type",
          v: ""
        },
        Owner: {
          k: "owner",
          v: ""
        },
        MemoRegexp: {
          k: "memo_regexp",
          v: ""
        }
      },
      type_18: {
        TxType: {
          k: "type",
          v: ""
        },
        From: {
          k: "from",
          v: "",
          f: "from_moniker"
        },
        Amount: {
          k: "amount",
          v: ""
        },
        To: {
          k: "to",
          v: "",
          f: "to_moniker"
        },
        "End Time": {
          k: "end-time",
          v: ""
        },
        from_moniker: {
          hide: true,
          k: "from_moniker",
          v: ""
        },
        to_moniker: {
          hide: true,
          k: "to_moniker",
          v: ""
        }
      },
      type_19: {
        TxType: {
          k: "type",
          v: ""
        },
        Amount: {
          k: "amount",
          v: ""
        },
        From: {
          k: "from",
          v: "",
          f: "from_moniker"
        },
        "Shares Src": {
          real_k: "Shares",
          k: "shares-src",
          v: ""
        },
        To: {
          k: "to",
          v: "",
          f: "to_moniker"
        },
        "Shares Dst": {
          real_k: "Shares",
          k: "shares-dst",
          v: ""
        },
        "End Time": {
          k: "end-time",
          v: ""
        },
        from_moniker: {
          hide: true,
          k: "from_moniker",
          v: ""
        },
        to_moniker: {
          hide: true,
          k: "to_moniker",
          v: ""
        }
      }
    };
  },
  computed: {
    txTypeSignData() {
      let data = this[this.txTypeSign];
      if (!data) {
        return [];
      }
      if (Array.isArray(data)) {
        return data;
      } else {
        return [data];
      }
    }
  },
  watch: {
    $route() {
      this.getTransactionInfo();
      Tools.scrollToTop();
    },
    "$store.state.isMobile"(newVal) {
      this.isMobileFunc(newVal);
    }
  },
  components: {
    MClip
  },
  beforeMount() {
    this.isMobileFunc(this.$store.state.isMobile);
  },
  mounted() {
    this.getTransactionInfo();
  },
  methods: {
    openUrl(url) {
      url = url.trim();
      if (url) {
        if (!/(http|https):\/\/([\w.]+\/?)\S*/.test(url)) {
          url = `http://${url}`;
        }
        window.open(url);
      }
    },
    isMobileFunc(isMobile) {
      if (isMobile) {
        this.transactionsDetailWrap = "mobile_transactions_detail_wrap";
      } else {
        this.transactionsDetailWrap =
          "personal_computer_transactions_detail_wrap";
      }
    },
    forAmount(data) {
      let amountValue = "";
      if (data.amount && data.amount.length !== 0) {
        amountValue = data.amount
          .map(item => {
            if (item.denom === Constant.Denom.IRISATTO) {
              return (item.amount = `${Tools.formatPriceToFixed(
                Tools.numberMoveDecimal(item.amount)
              )} ${Tools.formatDenom(item.denom).toLocaleUpperCase()}`);
            } else if (
              item.denom !== Constant.Denom.IRISATTO &&
              item.denom !== ""
            ) {
              return (item.amount = `${Tools.FormatScientificNotationToNumber(
                item.amount
              )} ${Tools.formatDenom(item.denom).toUpperCase()}`);
            } else {
              if (
                data.type === "BeginUnbonding" ||
                data.type === "BeginRedelegate"
              ) {
                return (item.amount = `${Tools.formatPriceToFixed(
                  Tools.numberMoveDecimal(item.amount)
                )} SHARES`);
              }
            }
          })
          .join(",");
      } else if (
        data.amount &&
        Object.keys(data.amount).includes("amount") &&
        Object.keys(data.amount).includes("denom")
      ) {
        data.amount = Tools.formatPriceToFixed(
          Tools.numberMoveDecimal(item.amount)
        );
        amountValue = `${data.amount.amount} ${Tools.formatDenom(
          data.amount.denom
        ).toUpperCase()}`;
      } else {
        amountValue = "";
      }
      return amountValue;
    },
    getKeyBaseName(identity, message) {
      let url = `https://keybase.io/_/api/1.0/user/lookup.json?fields=basics&key_suffix=${identity}`;
      if (identity) {
        axios.http(url).then(res => {
          if (res.them && res.them[0].basics.username) {
            message.identityUrl = `https://keybase.io/${res.them[0].basics.username}`;
          }
        });
      }
    },
    forMsgType(data) {
      if (!Array.isArray(data.msgs)) {
        return;
      }
      let msgs = data.msgs.map(v => {
        v.msg.type = v.type;
        return v.msg;
      });
      let arr = [];
      for (let it of msgs) {
        let o = JSON.parse(JSON.stringify(this[this.txTypeSign]));
        for (let i in o) {
          let fieidValue = it[o[i].k];
          if (this.typeValue !== "MintToken" && i === "Amount") {
            fieidValue = `${this.$options.filters.amountFromat(
              it[o[i].k],
              Constants.Denom.IRIS.toUpperCase()
            )}`;
          }
          if (i === "Identity") {
            this.getKeyBaseName(it[o[i].k], o[i]);
          }
          o[i].v =
            fieidValue === "" || fieidValue === undefined ? o[i].v : fieidValue;
        }
        arr.push(o);
      }
      this[this.txTypeSign] = arr;
    },
    getTransactionInfo() {
      if (this.$route.query.txHash) {
        Service.commonInterface(
          { txDetail: { txHash: this.$route.query.txHash } },
          data => {
            try {
              if (data) {
                // Transaction Information
                let that = this;
                let currentServerTime =
                  new Date().getTime() + that.diffMilliseconds;
                this.info.TxHash = data.hash;
                this.info.Block = data.block_height;
                this.info.Status = data.status
                  ? Tools.firstWordUpperCase(data.status)
                  : "";
                this.failInfo = data.log;
                this.ageValue = Tools.formatAge(
                  currentServerTime,
                  data.timestamp,
                  Constant.SUFFIX
                );
                clearInterval(this.transactionDetailTimer);
                this.transactionDetailTimer = setInterval(function() {
                  currentServerTime =
                    new Date().getTime() + that.diffMilliseconds;
                  that.ageValue = Tools.formatAge(
                    currentServerTime,
                    data.timestamp,
                    Constant.SUFFIX
                  );
                }, 1000);
                this.info.Timestamp = Tools.format2UTC(data.timestamp);
                this.info.Fee = `${Tools.convertScientificNotation2Number(
                  Tools.formatNumber(data.fee.amount)
                )} ${Tools.formatDenom(data.fee.denom).toUpperCase()}`;
                this.gasLimit = data.gas_limit;
                this.info.GasUsed = data.gas_used;
                this.gasWanted = data.gas_wanted;
                this.gasPrice = Tools.convertScientificNotation2Number(
                  Tools.formaNumberAboutGasPrice(data.gas_price)
                );
                this.info.Signer = data.signer || "";
                this.info.Memo = data.memo || "";
                // Transaction Message
                this.typeValue = data.type === "coin" ? "Transfer" : data.type;
                for (let [k, it] of Object.entries(this.txTypeArr)) {
                  if (it.includes(this.typeValue)) {
                    this.txTypeSign = k;
                    break;
                  }
                }
                if (
                  this.typeValue === "RequestRand" ||
                  this.typeValue === "SubmitProposal" ||
                  this.typeValue === "BeginUnbonding" ||
                  this.typeValue === "BeginRedelegate"
                ) {
                  data = Object.assign(data, data.tags);
                }
                if (
                  this.typeValue === "CreateValidator" ||
                  this.typeValue === "EditValidator"
                ) {
                  data = Object.assign(data, data.commission);
                }
                let message = this[this.txTypeSign];
                if (
                  this.typeValue === "SubmitProposal" &&
                  this[data.proposal_type]
                ) {
                  message = Object.assign(message, this[data.proposal_type]);
                }
                for (let i in message) {
                  let fieidValue = "";
                  if (i === "Amount") {
                    fieidValue = this.forAmount(data);
                  } else if (i === "Self-Bonded") {
                    if (data.self_bond && data.self_bond.length !== 0) {
                      fieidValue = `${Tools.formatPriceToFixed(
                        Tools.convertScientificNotation2Number(
                          Tools.formatNumber(data.self_bond[0].amount)
                        )
                      )} ${Tools.formatDenom(
                        data.self_bond[0].denom
                      ).toUpperCase()}`;
                    } else {
                      fieidValue = "";
                    }
                  } else if (i === "Initial Deposit" || i === "Deposit") {
                    if (data.amount && data.amount.length !== 0) {
                      fieidValue = this.$options.filters.amountFromat(
                        data.amount
                      );
                    } else {
                      fieidValue = "";
                    }
                  } else if (
                    (i === "Treshold" ||
                      i === "Commission Rate" ||
                      i === "Max Change Rate" ||
                      i === "Max Rate") &&
                    data[message[i].k] !== ""
                  ) {
                    fieidValue = `${data[message[i].k] * 100} %`;
                  } else if (
                    (this.typeValue === "WithdrawDelegatorRewardsAll" ||
                      this.typeValue === "WithdrawValidatorRewardsAll") &&
                    i === "From"
                  ) {
                    let arr = data[message[i].k].split(",");
                    fieidValue = arr.length === 1 ? arr[0] : arr;
                  } else if (i === "End Time") {
                    fieidValue = Tools.format2UTC(data[message[i].k]);
                  } else if (i === "Shares Src" || i === "Shares Dst") {
                    fieidValue = this.$options.filters.amountFromat({
                      amount: data[message[i].k]
                    });
                  } else {
                    fieidValue = data[message[i].k];
                  }
                  message[i].v = fieidValue;
                  if (i === "Identity") {
                    this.getKeyBaseName(data[message[i].k], message[i]);
                  }
                }
                if (
                  //asset类型交易
                  this.typeValue === "IssueToken" ||
                  this.typeValue === "EditToken" ||
                  this.typeValue === "MintToken" ||
                  this.typeValue === "TransferTokenOwner" ||
                  this.typeValue === "CreateGateway" ||
                  this.typeValue === "EditGateway" ||
                  this.typeValue === "TransferGatewayOwner" ||
                  this.typeValue === "RequestRand" ||
                  this.typeValue === "SetMemoRegexp"
                ) {
                  this.forMsgType(data);
                }
              }
            } catch (e) {
              console.error(e);
            }
          }
        );
      }
    }
  }
};
</script>
<style lang="scss" scoped>
@import "../style/mixin.scss";
.information_pre {
  color: #a2a2ae;
  white-space: pre-wrap;
}
.transactions_detail_wrap {
  @include flex;
  @include pcContainer;
  font-size: 0.14rem;
  .transactions_title_wrap {
    width: 100%;
    border-bottom: 1px solid #d6d9e0 !important;
    height: 0.62rem;
    background: #efeff1;
    line-height: 0.62rem;
    @include flex;
    @include pcContainer;
    .personal_computer_transactions_detail_wrap {
      height: 0.62rem;
      @include flex;
      span {
        line-height: 0.62rem;
        height: 0.62rem;
      }
    }
    .mobile_transactions_detail_wrap {
      @include flex;
      flex-direction: column;
    }
  }
  .personal_computer_transactions_detail_wrap {
    .transaction_information_content_title {
      height: 0.5rem;
      line-height: 0.5rem;
      font-size: 0.18rem;
      color: #000000;
      @include fontWeight;
      margin-left: 0.2rem;
      margin-bottom: 0;
      font-family: ArialMT;
    }
    @include pcCenter;
    .transactions_detail_information_wrap {
      padding: 0.2rem 0.2rem 0.08rem;
      border: 1px solid #d7d9e0;
      .information_props_wrap {
        @include flex;
        line-height: 0.2rem;
        margin-bottom: 0.12rem;
        .information_props {
          width: 1.5rem;
          color: var(--contentColor);
        }
        .information_value {
          color: var(--titleColor);
          flex: 1;
        }
      }
      &:nth-last-of-type(1) {
        margin-bottom: 0.4rem;
      }
    }
    .transactions_detail_title {
      height: 0.4rem;
      line-height: 0.4rem;
      font-size: 0.22rem;
      color: #000000;
      margin-right: 0.2rem;
      @include fontWeight;
    }
    .transactions_detail_wrap_hash_var {
      height: 0.4rem;
      line-height: 0.4rem;
      font-size: 0.22rem;
      color: #a2a2ae;
    }
  }
  .mobile_transactions_detail_wrap {
    width: 100%;
    @include flex;
    flex-direction: column;
    padding: 0 0.1rem;
    .transaction_information_content_title {
      height: 0.5rem;
      line-height: 0.5rem;
      font-size: 0.18rem;
      color: #000000;
      margin-left: 0.1rem;
      margin-bottom: 0;
      @include fontWeight;
    }
    .transactions_detail_information_wrap {
      padding: 0.1rem;
      border: 1px solid #d7d9e0;
      .information_props_wrap {
        @include flex;
        flex-direction: column;
        margin-bottom: 0;
        .information_value {
          overflow-x: auto;
          -webkit-overflow-scrolling: touch;
          overflow-y: hidden;
          color: var(--titleColor);
        }
      }
      &:nth-last-of-type(1) {
        margin-bottom: 0.4rem;
      }
    }
    .transactions_detail_title {
      height: 0.3rem;
      line-height: 0.3rem;
      font-size: 0.22rem;
      color: #000000;
      margin-right: 0.2rem;
      @include fontWeight;
    }
    .transactions_detail_wrap_hash_var {
      overflow-x: auto;
      -webkit-overflow-scrolling: touch;
      height: 0.3rem;
      line-height: 0.3rem;
      font-size: 0.18rem;
      color: #a2a2ae;
    }
  }
}
.link_active_style {
  a {
    color: var(--bgColor) !important;
  }
  color: var(--bgColor) !important;
  cursor: pointer;
}
.information_value_fixed {
  display: flex;
  align-items: center;
  & > span {
    margin-right: 0.06rem;
  }
  .fail_status {
    color: #fa8593;
  }
  .question_icon_div {
    background-image: url(../assets/question_icon.png) !important;
  }
  .info_icon_div {
    width: 0.14rem;
    height: 0.14rem;
    position: relative;
    background: url(../assets/info_icon.png) no-repeat top left / 14px 14px;
    cursor: pointer;
    &:hover {
      .tooltip_span {
        display: block;
        position: fixed;
        opacity: 0;
      }
    }
    .tooltip_span {
      display: none;
      z-index: 1000;
      color: #ffffff;
      background-color: #000000;
      border-radius: 0.04rem;
      line-height: 16px;
      div {
        padding: 8px 15px;
        & > p {
          white-space: nowrap;
        }
      }
      & > i {
        width: 0;
        height: 0;
        border: 0.06rem solid transparent;
        content: "";
        display: block;
        position: absolute;
        border-top-color: #000000;
        margin-left: -6px;
      }
    }
    .tooltip_span_word_warp {
      word-break: break-all;
      word-wrap: break-word;
      white-space: normal;
    }
  }
}
</style>
