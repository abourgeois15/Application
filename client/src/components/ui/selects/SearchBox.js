import React from "react";
import PropTypes from 'prop-types';

const SearchBox = (props) => {
  return (
    <div data-cy="item-search-box">
      <input
        className="inputBox"
        placeholder="Search Items"
        onChange={props.handleChange}
        value={props.value}
      />
    </div>
  );
};

SearchBox.propTypes = {
  handleChange: PropTypes.func,
  value: PropTypes.string
};

export default SearchBox;