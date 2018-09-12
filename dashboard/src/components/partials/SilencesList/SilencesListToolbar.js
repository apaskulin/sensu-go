import React from "react";
import PropTypes from "prop-types";

import ListToolbar from "/components/partials/ListToolbar";
import NewMenuItem from "/components/partials/ToolbarMenuItems/New";
import ResetMenuItem from "/components/partials/ToolbarMenuItems/Reset";
import SearchBox from "/components/SearchBox";
import ToolbarMenu from "/components/partials/ToolbarMenu";

class SilencesListToolbar extends React.Component {
  static propTypes = {
    filter: PropTypes.string,
    onChangeQuery: PropTypes.func.isRequired,
    onClickCreate: PropTypes.func.isRequired,
    onClickReset: PropTypes.func.isRequired,
  };

  static defaultProps = {
    filter: "",
  };

  render() {
    return (
      <ListToolbar
        search={
          <SearchBox
            placeholder="Filter silences…"
            initialValue={this.props.filter}
            onSearch={this.props.onChangeQuery}
          />
        }
        toolbarItems={({ collapsed }) => (
          <ToolbarMenu>
            <ToolbarMenu.Item visible={collapsed ? "never" : "always"}>
              <NewMenuItem
                title="New Silence…"
                onClick={this.props.onClickCreate}
              />
            </ToolbarMenu.Item>
            <ToolbarMenu.Item visible={collapsed ? "never" : "if-room"}>
              <ResetMenuItem onClick={this.props.onClickReset} />
            </ToolbarMenu.Item>
          </ToolbarMenu>
        )}
      />
    );
  }
}

export default SilencesListToolbar;
