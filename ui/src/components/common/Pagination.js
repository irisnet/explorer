import React, { Component } from "react"
import { Menu } from "semantic-ui-react"
import "./Pagination.css"

export default class Pagination extends Component {
  state = { activeItem: this.props.currentPage }

  handleItemClick = (e, { name }) => this.setState({ activeItem: name })

  render() {
    const { activeItem } = this.state
    const { totalPagesCount } = this.props
    if (totalPagesCount <= 1) return <div />
    const range = Array.from({ length: totalPagesCount }, (x, i) => (i + 1).toString())

    return (
      <Menu pagination>
        {range.map(r => (
          <Menu.Item key={r} name={r} active={activeItem === r} onClick={this.handleItemClick} />
        ))}
      </Menu>
    )
  }
}
