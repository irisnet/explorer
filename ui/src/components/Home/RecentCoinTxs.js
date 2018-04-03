import React, { Component } from "react"
import { Segment, Table, Header } from "semantic-ui-react"
import { inject, observer } from "mobx-react"
import { Link } from "react-router-dom"

import RedError from "../common/RedError"
import { timeAgo } from "../../utils/TimeAgo"

@inject("txStore")
@observer
class RecentCoinTxs extends Component {
  componentDidMount() {
    this.props.txStore.loadRecentCoinTx()
  }
  render() {
    const { isLoading, error } = this.props.txStore
    if (error !== undefined) return <RedError message={error} />

    const recent = this.props.txStore.recentCoinTx || []
    return (
      <Segment basic loading={isLoading}>
        <Header>Coin Transactions</Header>
        <Table compact fixed singleLine>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Hash</Table.HeaderCell>
              <Table.HeaderCell width={5}>Time</Table.HeaderCell>
              <Table.HeaderCell>From</Table.HeaderCell>
              <Table.HeaderCell>To</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {recent.slice(0, 20).map((v, index) => {
              return (
                <Table.Row key={index} verticalAlign="top">
                  <Table.Cell>
                    <Link to={"/tx/" + v.tx_hash}>{v.tx_hash}</Link>
                  </Table.Cell>
                  <Table.Cell>{timeAgo(v.time)}</Table.Cell>
                  <Table.Cell>
                    <Link to={"/account/" + v.from}>{v.from}</Link>
                  </Table.Cell>
                  <Table.Cell>
                    <Link to={"/account/" + v.to}>{v.to}</Link>
                  </Table.Cell>
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>
      </Segment>
    )
  }
}

export default RecentCoinTxs
