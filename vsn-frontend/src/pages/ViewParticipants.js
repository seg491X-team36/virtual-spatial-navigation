import {useState} from 'react'
import 'semantic-ui-css/semantic.min.css'
import { Container, Header, Grid, Menu } from 'semantic-ui-react';

import * as fake_data from '../components/apiFunctions.js'
import ParticipantDetails from '../components/ParticipantDetails.js'

const ViewParticipants = () => {
    const participants = fake_data.users('ACCEPTED');
    const[selectedUser, setSelectedUser] = useState({});

    const handleViewUserClick = (u) => {
        setSelectedUser({
            id: u.id,
            name: u.name,
            description: u.description,
            config: u.config,
            results: u.results,
            pending: u.pending,
            usersNotInvited: u.usersNotInvited
        });
    }

    return(
        <Container>
        <Header as='h1'>Participants</Header>
        <Grid columns='equal'>
            <Grid.Column>
                <Menu vertical>
                {participants.length > 0 ? participants.map(p => {return (
                <Menu.Item
                    name={p.email}
                    onClick={() => handleViewUserClick(p)}
                    active={selectedUser.id === p.id}>
                    <Header as='h4'>{p.email}</Header>
                </Menu.Item>
                )})
                :
                <p>There are no participants to display.</p>
                }
                </Menu>
            </Grid.Column>

            <Grid.Column width={12}>
            {selectedUser.id !== undefined ?
                <ParticipantDetails user_id={selectedUser.id}/>
                :
                <></>
                }
            </Grid.Column>
        </Grid>
        </Container>
    )
};

export default ViewParticipants;