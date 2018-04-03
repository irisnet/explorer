import React, { Component } from "react"
import { Grid } from "semantic-ui-react"

import RecentBlocks from "./RecentBlocks"
import RecentCoinTxs from "./RecentCoinTxs"
import RecentStakeTxs from "./RecentStakeTxs"

export default class Home extends Component {
  render() {
    return (
      <Grid columns={3}>
        <Grid.Column style={{ padding: 0 }}>
          <RecentBlocks />
        </Grid.Column>
        <Grid.Column style={{ padding: 0 }}>
          <RecentCoinTxs />
        </Grid.Column>
        <Grid.Column style={{ padding: 0 }}>
          <RecentStakeTxs />
        </Grid.Column>
      </Grid>
    )
  }
}
