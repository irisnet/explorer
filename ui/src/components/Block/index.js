import React, { Component } from "react"
import { inject, observer } from "mobx-react"
import { Segment, Header, Table, Button, Icon } from "semantic-ui-react"
import { Link } from "react-router-dom"

import Validators from "./Validators"
import RedError from "../common/RedError"

@inject("blockStore")
@inject("commonStore")
@observer
export default class Block extends Component {
  componentDidMount() {
    const height = this.props.match.params.height
    this.props.blockStore.loadBlock(height)
  }
  componentWillReceiveProps(nextProps) {
    var h_old = this.props.match.params.height
    var h_new = nextProps.match.params.height
    if (h_old !== h_new) {
      this.props.blockStore.loadBlock(h_new, { acceptCached: true })
    }
  }

  render() {
    const { isLoading, error } = this.props.blockStore
    if (error !== undefined) return <RedError message={error} />

    const height = this.props.match.params.height
    const block = this.props.blockStore.getBlock(height)
    if (!block) return <div />

    const meta = block.block_meta
    const prev = height - 1
    const next = height - 1 + 2
    const { status } = this.props.commonStore
    const latest_block_height = status.latest_block_height

    return (
      <div>
        <Header dividing>Details for Block #{height}</Header>
        <Segment basic compact loading={isLoading}>
          <Table basic="very" fixed singleLine>
            <Table.Body>
              <Table.Row>
                <Table.Cell width={1}>Hash</Table.Cell>
                <Table.Cell width={10}>
                  <Button
                    icon
                    compact
                    basic
                    circular
                    as={Link}
                    to={"/block/" + prev}
                    disabled={prev <= 0}
                  >
                    <Icon name="chevron left" color="blue" />
                  </Button>
                  {meta.block_id.hash}
                  <Button
                    icon
                    compact
                    basic
                    circular
                    as={Link}
                    to={"/block/" + next}
                    disabled={next > latest_block_height}
                  >
                    <Icon name="chevron right" color="blue" />
                  </Button>
                </Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell width={1}>Time</Table.Cell>
                <Table.Cell width={10}>{meta.header.time}</Table.Cell>
              </Table.Row>
              <Table.Row>
                <Table.Cell width={1}>Transactions</Table.Cell>
                <Table.Cell width={10}>
                  <NumTxs num_txs={meta.header.num_txs} data_hash={meta.header.data_hash} />
                </Table.Cell>
              </Table.Row>
            </Table.Body>
          </Table>
        </Segment>
        <Header dividing>Validators</Header>
        <Validators height={height} />
      </div>
    )
  }
}

const NumTxs = ({ num_txs, data_hash }) =>
  num_txs > 0 ? <Link to={"/tx/" + data_hash}>{num_txs}</Link> : <span>{num_txs}</span>
