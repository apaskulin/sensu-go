import React from "react";
import PropTypes from "prop-types";
import gql from "graphql-tag";

import Content from "/components/Content";
import Grid from "@material-ui/core/Grid";
import Loader from "/components/util/Loader";
import RelatedEntitiesCard from "/components/partials/RelatedEntitiesCard";

import CheckResult from "./EventDetailsCheckResult";
import Toolbar from "./EventDetailsToolbar";
import Summary from "./EventDetailsSummary";

class EventDetailsContainer extends React.PureComponent {
  static propTypes = {
    event: PropTypes.object,
    loading: PropTypes.bool.isRequired,
    poller: PropTypes.shape({
      running: PropTypes.bool,
      start: PropTypes.func,
      stop: PropTypes.func,
    }).isRequired,
  };

  static defaultProps = {
    event: null,
  };

  static fragments = {
    event: gql`
      fragment EventDetailsContainer_event on Event {
        id
        timestamp
        deleted @client
        ...EventDetailsToolbar_event

        check {
          ...EventDetailsCheckResult_check
          ...EventDetailsSummary_check
        }
        entity {
          ...EventDetailsCheckResult_entity
          ...RelatedEntitiesCard_entity
          ...EventDetailsSummary_entity
        }
      }

      ${CheckResult.fragments.check}
      ${CheckResult.fragments.entity}
      ${RelatedEntitiesCard.fragments.entity}
      ${Summary.fragments.check}
      ${Summary.fragments.entity}
      ${Toolbar.fragments.event}
    `,
  };

  render() {
    const { event, loading, poller } = this.props;

    return (
      <Loader loading={loading} passthrough>
        {event && (
          <React.Fragment>
            <Content marginBottom>
              <Toolbar event={event} poller={poller} />
            </Content>
            <Grid container spacing={16}>
              <Grid item xs={12}>
                <CheckResult check={event.check} entity={event.entity} />
              </Grid>
              <Grid item xs={12} md={6}>
                <RelatedEntitiesCard entity={event.entity} />
              </Grid>
              <Grid item xs={12} md={6}>
                <Summary check={event.check} entity={event.entity} />
              </Grid>
            </Grid>
          </React.Fragment>
        )}
      </Loader>
    );
  }
}

export default EventDetailsContainer;
