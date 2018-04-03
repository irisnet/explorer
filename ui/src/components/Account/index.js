import React, { Component } from "react"
import { inject, observer } from "mobx-react"
import { Segment, Header, Table } from "semantic-ui-react"

import RedError from "../common/RedError"
import Coins from "./Coins"
import CoinTxs from "./CoinTxs"

@inject("accountStore")
@observer
export default class Account extends Component {
  componentDidMount() {
    const address = this.props.match.params.address
    this.props.accountStore.loadAccount(address)
  }
  componentWillReceiveProps(nextProps) {
    var addr_old = this.props.match.params.address
    var addr_new = nextProps.match.params.address
    if (addr_old !== addr_new) {
      this.props.accountStore.loadAccount(addr_new, { acceptCached: true })
    }
  }

  render() {
    const { isLoading, error } = this.props.accountStore
    if (error !== undefined) return <RedError message={error} />

    const address = this.props.match.params.address
    const account = this.props.accountStore.getAccount(address) || {}

    return (
      <div>
        <Header dividing>Details for Address</Header>
        <Segment basic compact loading={isLoading}>
          <Table basic="very" fixed singleLine>
            <Table.Body>
              <Table.Row>
                <Table.Cell width={1}>Address</Table.Cell>
                <Table.Cell width={10}>{address}</Table.Cell>
              </Table.Row>
              <Table.Row verticalAlign="top">
                <Table.Cell>Coins</Table.Cell>
                <Table.Cell>
                  <Coins coins={account.coins} />
                </Table.Cell>
              </Table.Row>
              <Table.Row verticalAlign="top">
                <Table.Cell>Credit</Table.Cell>
                <Table.Cell>
                  <Coins coins={account.credit} />
                </Table.Cell>
              </Table.Row>
            </Table.Body>
          </Table>
        </Segment>
        <Header dividing>Coin Transactions</Header>
        <CoinTxs address={address} />
      </div>
    )
  }
}
