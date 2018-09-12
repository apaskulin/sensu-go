import React from "react";
import PropTypes from "prop-types";
import gql from "graphql-tag";

import Checkbox from "@material-ui/core/Checkbox";
import TableCell from "@material-ui/core/TableCell";

import ConfirmDelete from "/components/partials/ConfirmDelete";
import DeleteMenuItem from "/components/partials/ToolbarMenuItems/Delete";
import SilenceMenuItem from "/components/partials/ToolbarMenuItems/Silence";
import UnsilenceMenuItem from "/components/partials/ToolbarMenuItems/Unsilence";
import ToolbarMenu from "/components/partials/ToolbarMenu";

import ResourceDetails from "/components/partials/ResourceDetails";
import TableOverflowCell from "/components/partials/TableOverflowCell";
import TableSelectableRow from "/components/partials/TableSelectableRow";

import { RelativeToCurrentDate } from "/components/RelativeDate";
import CheckStatusIcon from "/components/CheckStatusIcon";
import NamespaceLink from "/components/util/NamespaceLink";

class EntitiesListItem extends React.PureComponent {
  static propTypes = {
    entity: PropTypes.object.isRequired,
    selected: PropTypes.bool,
    onChangeSelected: PropTypes.func,
    onClickClearSilence: PropTypes.func,
    onClickDelete: PropTypes.func,
    onClickSilence: PropTypes.func,
  };

  static defaultProps = {
    selected: undefined,
    onChangeSelected: ev => ev,
    onClickClearSilence: ev => ev,
    onClickDelete: ev => ev,
    onClickSilence: ev => ev,
  };

  static fragments = {
    entity: gql`
      fragment EntitiesListItem_entity on Entity {
        id
        name
        lastSeen
        class
        status
        isSilenced
        system {
          platform
          platformVersion
        }
      }
    `,
  };

  renderDescription = () => {
    const { entity } = this.props;

    return (
      <React.Fragment>
        <strong>{entity.class}</strong> - Last seen{" "}
        {entity.lastSeen && (
          <strong>
            <RelativeToCurrentDate dateTime={entity.lastSeen} />{" "}
          </strong>
        )}
        with status <strong>{entity.status}</strong>.
      </React.Fragment>
    );
  };

  render() {
    const { entity, selected, onChangeSelected } = this.props;

    return (
      <TableSelectableRow selected={selected}>
        <TableCell padding="checkbox">
          <Checkbox
            color="primary"
            checked={selected}
            onChange={e => onChangeSelected(e.target.checked)}
          />
        </TableCell>

        <TableOverflowCell>
          <ResourceDetails
            icon={
              <CheckStatusIcon
                statusCode={entity.status}
                silenced={entity.isSilenced}
              />
            }
            title={
              <NamespaceLink
                namespace={entity.namespace}
                to={`/entities/${entity.name}`}
              >
                <strong>{entity.name}</strong> {entity.system.platform}{" "}
                {entity.system.platformVersion}
              </NamespaceLink>
            }
            details={this.renderDescription()}
          />
        </TableOverflowCell>

        <TableCell padding="checkbox">
          <ToolbarMenu>
            <ToolbarMenu.Item visible="never">
              <SilenceMenuItem
                disabled={entity.isSilenced}
                onClick={this.props.onClickSilence}
              />
            </ToolbarMenu.Item>
            <ToolbarMenu.Item visible="never">
              <UnsilenceMenuItem
                disabled={!entity.isSilenced}
                onClick={this.props.onClickClearSilence}
              />
            </ToolbarMenu.Item>
            <ToolbarMenu.Item visible="never">
              {menu => (
                <ConfirmDelete
                  onSubmit={() => {
                    this.props.onClickDelete();
                    menu.close();
                  }}
                >
                  {dialog => (
                    <DeleteMenuItem
                      autoClose={false}
                      title="Delete…"
                      onClick={dialog.open}
                    />
                  )}
                </ConfirmDelete>
              )}
            </ToolbarMenu.Item>
          </ToolbarMenu>
        </TableCell>
      </TableSelectableRow>
    );
  }
}

export default EntitiesListItem;
