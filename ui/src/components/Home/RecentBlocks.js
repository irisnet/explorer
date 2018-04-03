import React, { Component } from "react"
import { Segment, Table, Header } from "semantic-ui-react"
import { inject, observer } from "mobx-react"
import { Link } from "react-router-dom"

import RedError from "../common/RedError"
import { timeAgo } from "../../utils/TimeAgo"

@inject("blockStore")
@observer
class RecentBlocks extends Component {
  componentDidMount() {
    this.props.blockStore.loadRecent()
  }
  render() {
    const { isLoading, error } = this.props.blockStore
    if (error !== undefined) return <RedError message={error} />

    const recent = this.props.blockStore.recent || []
    return (
      <Segment basic loading={isLoading}>
        <Header>Recent Blocks</Header>
        <Table compact fixed singleLine>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell width={1}>Height</Table.HeaderCell>
              <Table.HeaderCell width={1}>Time</Table.HeaderCell>
              <Table.HeaderCell width={1}>Txns</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {recent.slice(0, 20).map((v, index) => {
              return (
                <Table.Row key={index} verticalAlign="top">
                  <Table.Cell>
                    <Link to={"/block/" + v.height}>{v.height}</Link>
                  </Table.Cell>
                  <Table.Cell>{timeAgo(v.time)}</Table.Cell>
                  <Table.Cell>{v.num_txs} </Table.Cell>
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>
      </Segment>
    )
  }
}

export default RecentBlocks
