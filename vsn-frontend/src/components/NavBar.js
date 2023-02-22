import React, {useState} from 'react';
import { Link, Outlet } from "react-router-dom";
import { Menu } from 'semantic-ui-react'

const NavBar = () => {
  return (
    <>
    <Menu>
      <Menu.Item name='home'>
        <Link to="/">Home</Link>
      </Menu.Item>

      <Menu.Item name='experiments'>
        <Link to="/experiments">View Experiments</Link>
      </Menu.Item>

      <Menu.Item name='participants'>
        <Link to="/participants">View Participants</Link>
      </Menu.Item>
      <Menu.Item name='rejected-sign-ups'>
        <Link to="/rejected-sign-ups">View Rejected Sign Ups</Link>
      </Menu.Item>

      <Menu.Item name='sign-ups'>
        <Link to="/sign-ups">View Sign Ups</Link>
      </Menu.Item>
    </Menu>

    <Outlet/>
    </>
  )
};

export default NavBar;