import React, { Component } from "react"
import { withRouter } from "react-router-dom"

import { Form } from "semantic-ui-react"

class SearchBox extends Component {
  constructor(props) {
    super(props)
    this.state = { value: "" }

    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
  }

  handleChange(event) {
    this.setState({ value: event.target.value })
  }

  handleSubmit(event) {
    event.preventDefault()
    this.setState({ value: "" })

    const value = this.state.value
    const path = isNaN(value) ? "/account/" + value : "/block/" + value
    this.props.history.push(path)
  }

  render() {
    return (
      <Form onSubmit={this.handleSubmit} size="mini">
        <Form.Input
          style={{ width: "350px" }}
          icon="search"
          value={this.state.value}
          onChange={this.handleChange}
          placeholder="Search for account, block..."
        />
      </Form>
    )
  }
}

export default withRouter(SearchBox)
