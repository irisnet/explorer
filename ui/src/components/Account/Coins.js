import React, { Component } from "react"
import { List } from "semantic-ui-react"

class Coins extends Component {
  render() {
    const coins = this.props.coins || []

    return (
      <List>
        {coins.map((coin, index) => (
          <List.Item key={index}>
            {coin.amount} {coin.denom}
          </List.Item>
        ))}
      </List>
    )
  }
}

export default Coins
