import React from "react"
import { Message } from "semantic-ui-react"

export default class RedError extends React.Component {
  render() {
    return <Message negative>{this.props.message}</Message>
  }
}
