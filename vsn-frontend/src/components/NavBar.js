import React, {useState} from 'react';
import { Link, Outlet } from "react-router-dom";
import 'semantic-ui-css/semantic.min.css'
import { Menu, Grid, Icon, Header, Segment, Sidebar } from 'semantic-ui-react'

const NavBar = (activatedLink) => {
  //const[visible, setVisible] = useState(false);
  const[activeItem, setActiveItem] = useState(activatedLink);

  return (
    <>
    <Menu pointing secondary>
      <Menu.Item
        name='Home'
        icon='home'
        active={activeItem === 'home'}
        onClick={()=>setActiveItem('home')}
        href='/'
      />
      <Menu.Item
        name='View Experiments'
        icon='chart bar'
        active={activeItem === 'experiments'}
        onClick={()=>setActiveItem('experiments')}
        href='/experiments'
      />
      <Menu.Item
        name='View Participants'
        icon='users'
        active={activeItem === 'participants'}
        onClick={()=>setActiveItem('participants')}
        href='/participants'
      />
      <Menu.Item
        name='View Sign Ups'
        icon='user plus'
        active={activeItem === 'signups'}
        onClick={()=>setActiveItem('signups')}
        href='/sign-ups'
      />
      <Menu.Item
        name='View Rejected Sign Ups'
        icon='user times'
        active={activeItem === 'rejected'}
        onClick={()=>setActiveItem('rejected')}
        href='/rejected-sign-ups'
      />
    </Menu>
    {/*
    <Icon name='content' size='large' onClick={() => setVisible(!visible)}/>
    <Grid columns={1}>
      <Grid.Column>
        <Sidebar.Pushable>
          <Sidebar
            as={Menu}
            animation='overlay'
            icon='labeled'
            inverted
            onHide={() => setVisible(false)}
            vertical
            visible={visible}
            width='thin'
          >
            <Menu.Item
              name='Home'
              icon='home'
              active={activeItem === 'home'}
              onClick={(e)=>setActiveItem('home')}
              href='/'
            />
            <Menu.Item
              name='View Experiments'
              icon='chart bar'
              active={activeItem === 'experiments'}
              onClick={(e)=>setActiveItem('experiments')}
              href='/experiments'
            />
            <Menu.Item
              name='View Participants'
              icon='users'
              active={activeItem === 'participants'}
              onClick={(e)=>setActiveItem('participants')}
              href='/participants'
            />
            <Menu.Item
              name='View Sign Ups'
              icon='user plus'
              active={activeItem === 'signups'}
              onClick={(e)=>setActiveItem('signups')}
              href='/sign-ups'
            />
            <Menu.Item
              name='View Rejected Sign Ups'
              icon='user times'
              active={activeItem === 'rejected'}
              onClick={(e)=>setActiveItem('rejected')}
              href='/rejected-sign-ups'
            />
          </Sidebar>

          <Sidebar.Pusher>
            <div height='100vh'>
              Application content
              Lorem ipsum dolor sit amet, consectetur adipiscing elit,
              sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
              Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
              nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
              reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
              pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
              culpa qui officia deserunt mollit anim id est laborum.
            </div>
          </Sidebar.Pusher>
        </Sidebar.Pushable>
      </Grid.Column>
    </Grid>
    */}

    <Outlet/>
    </>
  )
};

export default NavBar;