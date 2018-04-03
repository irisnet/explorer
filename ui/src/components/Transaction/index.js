import React, { Component } from "react"
import { inject, observer } from "mobx-react"
import { Segment, Header, Table } from "semantic-ui-react"
import { Link } from "react-router-dom"

import RedError from "../common/RedError"

@inject("txStore")
@inject("commonStore")
@observer
export default class Transaction extends Component {
  componentDidMount() {
    const txhash = this.props.match.params.txhash
    this.props.txStore.loadTx(txhash)
  }
  componentWillReceiveProps(nextProps) {
    var h_old = this.props.match.params.txhash
    var h_new = nextProps.match.params.txhash
    if (h_old !== h_new) {
      this.props.txStore.loadTx(h_new, { acceptCached: true })
    }
  }

  render() {
    const { isLoading, error } = this.props.txStore
    if (error !== undefined) return <RedError message={error} />

    const txhash = this.props.match.params.txhash
    const tx = this.props.txStore.getTx(txhash)
    if (!tx) return <div />

    return (
      <div>
        <Header dividing>Details for Transaction</Header>
        <Segment basic compact loading={isLoading}>
          <Table basic="very" fixed singleLine>
            <Table.Body>
              <Table.Row>
                <Table.Cell width={1}>Hash</Table.Cell>
                <Table.Cell width={10}>{txhash}</Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell width={1}>Height</Table.Cell>
                <Table.Cell width={10}>
                  <Link to={"/block/" + tx.height}>{tx.height}</Link>
                </Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell width={1}>Time</Table.Cell>
                <Table.Cell width={10}>{tx.time}</Table.Cell>
              </Table.Row>
            </Table.Body>
          </Table>
          <Segment>
            <pre>{JSON.stringify(tx, null, 2)}</pre>
          </Segment>
        </Segment>
      </div>
    )
  }
}
