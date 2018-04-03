import Header from "./Header"
import React from "react"
import { Switch, Route, withRouter } from "react-router-dom"
import { inject, observer } from "mobx-react"
import { Container, Segment, Label } from "semantic-ui-react"
import { Link } from "react-router-dom"

import Home from "./Home"
import Account from "./Account"
import Block from "./Block"
import Transaction from "./Transaction"

@inject("commonStore")
@withRouter
@observer
export default class App extends React.Component {
  componentDidMount() {
    this.props.commonStore.loadStatus()
  }
  componentWillReceiveProps(nextProps) {
    this.props.commonStore.loadStatus()
  }

  render() {
    const { isLoading, status } = this.props.commonStore
    return (
      <div>
        <Header />
        <Container>
          <Segment basic compact loading={isLoading} style={{ marginTop: "5em" }}>
            <Label>
              LAST BLOCK:
              <Label.Detail as={Link} to={"/block/" + status.last_height}>
                {status.last_height}
              </Label.Detail>
            </Label>
          </Segment>
          <Switch>
            <Route path="/account/:address" component={Account} />
            <Route path="/block/:height" component={Block} />
            <Route path="/tx/:txhash" component={Transaction} />
            <Route path="/" component={Home} />
          </Switch>
        </Container>
      </div>
    )
  }
}
