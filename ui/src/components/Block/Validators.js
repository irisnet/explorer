import React, { Component } from "react"
import { Segment, Table } from "semantic-ui-react"
import { inject, observer } from "mobx-react"

import RedError from "../common/RedError"

@inject("validatorStore")
@observer
class Validators extends Component {
  componentDidMount() {
    const height = this.props.height
    this.props.validatorStore.loadValidators(height)
  }
  componentWillReceiveProps(nextProps) {
    var height_old = this.props.address
    var height_new = nextProps.address
    if (height_old !== height_new) {
      this.props.validatorStore.loadValidators(height_new)
    }
  }
  render() {
    const { isLoading, error } = this.props.validatorStore
    if (error !== undefined) return <RedError message={error} />

    const validators = this.props.validatorStore.validators || []
    return (
      <Segment basic loading={isLoading}>
        <Table basic="very" compact>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Address</Table.HeaderCell>
              <Table.HeaderCell>Voting Power</Table.HeaderCell>
              <Table.HeaderCell>Accum</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {validators.map((v, index) => <Validator key={index} validator={v} />)}
          </Table.Body>
        </Table>
      </Segment>
    )
  }
}

const Validator = ({ validator }) => {
  const { address, voting_power, accum } = validator
  return (
    <Table.Row verticalAlign="top">
      <Table.Cell>{address} </Table.Cell>
      <Table.Cell> {voting_power} </Table.Cell>
      <Table.Cell> {accum} </Table.Cell>
    </Table.Row>
  )
}

export default Validators
