import React, { Component } from "react"
import { Segment, Table, Label } from "semantic-ui-react"
import { inject, observer } from "mobx-react"
import { Link } from "react-router-dom"

import Coins from "./Coins"
import RedError from "../common/RedError"
import Pagination from "../common/Pagination"

@inject("coinTxsStore")
@observer
class CoinTxs extends Component {
  componentDidMount() {
    const address = this.props.address
    this.props.coinTxsStore.loadCoinTxs(address)
  }
  componentWillReceiveProps(nextProps) {
    var addr_old = this.props.address
    var addr_new = nextProps.address
    if (addr_old !== addr_new) {
      this.props.coinTxsStore.loadCoinTxs(addr_new)
    }
  }
  render() {
    const { isLoading, error } = this.props.coinTxsStore
    if (error !== undefined) return <RedError message={error} />

    const txs = this.props.coinTxsStore.coinTxs || []
    return (
      <Segment basic loading={isLoading}>
        <Table basic="very" compact>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell width={1}>Height</Table.HeaderCell>
              <Table.HeaderCell width={5}>From</Table.HeaderCell>
              <Table.HeaderCell width={1} />
              <Table.HeaderCell width={5}>To</Table.HeaderCell>
              <Table.HeaderCell width={1}>TxFee</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {txs.map((tx, index) => <CoinTx key={index} height={tx.height} tx={tx.tx} />)}
          </Table.Body>
        </Table>
        <Pagination totalPagesCount="1" currentPage="1" />
      </Segment>
    )
  }
}

@inject("coinTxsStore")
class CoinTx extends Component {
  getCoinLabel(inputs, outputs, address) {
    let isFrom = inputs.some(o => o.address.addr === address)
    let isTo = outputs.some(o => o.address.addr === address)
    return isFrom && isTo ? "SELF" : isFrom ? "OUT" : isTo ? "IN" : ""
  }
  parseTx(tx) {
    let fee = []
    if (tx.type === "fee/tx") {
      fee.push(tx.data.fee)
      tx = tx.data.tx.data
    } else {
      tx = tx.data
    }
    return { fee, tx }
  }

  render() {
    const height = this.props.height
    const address = this.props.coinTxsStore.address
    let { tx, fee } = this.parseTx(this.props.tx)
    const { inputs, outputs } = tx
    const label = this.getCoinLabel(inputs, outputs, address)

    return (
      <Table.Row verticalAlign="top">
        <Table.Cell>
          <Link to={"/block/" + height}>{height}</Link>
        </Table.Cell>
        <Table.Cell>
          {inputs.map((input, index) => (
            <InOut key={index} address={input.address.addr} coins={input.coins} />
          ))}
        </Table.Cell>
        <Table.Cell>
          <CoinLabel label={label} />
        </Table.Cell>
        <Table.Cell>
          {outputs.map((output, index) => (
            <InOut key={index} address={output.address.addr} coins={output.coins} />
          ))}
        </Table.Cell>
        <Table.Cell>
          <Coins coins={fee} />
        </Table.Cell>
      </Table.Row>
    )
  }
}

@inject("coinTxsStore")
class InOut extends Component {
  render() {
    const { address, coins } = this.props
    const me = this.props.coinTxsStore.address === address

    return (
      <Table basic="very" compact fixed singleLine>
        <Table.Body>
          <Table.Row verticalAlign="top">
            <Table.Cell width={3} style={{ paddingTop: "0px" }}>
              <Address address={address} isLink={!me} />
            </Table.Cell>
            <Table.Cell width={1} style={{ paddingTop: "0px" }}>
              <Coins coins={coins} />
            </Table.Cell>
          </Table.Row>
        </Table.Body>
      </Table>
    )
  }
}

const Address = ({ address, isLink }) =>
  isLink ? <Link to={"/account/" + address}>{address}</Link> : <span>{address}</span>

const CoinLabel = ({ label }) => {
  const color =
    label === "IN" ? "green" : label === "OUT" ? "orange" : label === "SELF" ? "blue" : ""
  if (color === "") return <div />
  return (
    <Label size="small" color={color} horizontal>
      {label}
    </Label>
  )
}

export default CoinTxs
