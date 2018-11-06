import React from "react";
import PropTypes from "prop-types";
import gql from "graphql-tag";
import capitalize from "lodash/capitalize";

import Maybe from "/components/Maybe";
import { RelativeToCurrentDate } from "/components/RelativeDate";

class EntityStatusDescriptor extends React.PureComponent {
  static propTypes = {
    entity: PropTypes.object.isRequired,
  };

  static fragments = {
    entity: gql`
      fragment EntityStatusDescriptor_entity on Entity {
        lastSeen
        class
      }
    `,
  };

  render() {
    const { entity } = this.props;

    if (!entity.lastSeen && entity.class !== "agent") {
      return (
        <React.Fragment>
          <strong>{capitalize(entity.class)}</strong> entity.
        </React.Fragment>
      );
    }

    return (
      <React.Fragment>
        The <strong>{entity.class}</strong> was last seen{" "}
        <strong>
          <Maybe value={entity.lastSeen} fallback="never">
            {val => <RelativeToCurrentDate direction="past" dateTime={val} />}
          </Maybe>
        </strong>.
      </React.Fragment>
    );
  }
}

export default EntityStatusDescriptor;
