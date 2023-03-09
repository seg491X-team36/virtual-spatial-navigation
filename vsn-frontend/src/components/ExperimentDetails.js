import React from 'react';
import 'semantic-ui-css/semantic.min.css'

import * as fake_data from '../components/apiFunctions.js';
import {Container, Header, List, Button, Grid, Checkbox, Segment} from 'semantic-ui-react';

const ExperimentDetails = ({exp_id}) => {

    let selected_exp = fake_data.experiment(exp_id);
    let selectedUsers = [];

    const handleCheckbox = (data) => {
        if (data.checked) {
            selectedUsers.push(data.value);
        } else {
            selectedUsers = selectedUsers.filter(u => u !== data.value);
        }
        console.log(selectedUsers);
    }

    const handleInviteUser = (user) => {
        console.log('handleInviteUser');
    }

    const handleInviteSelected = () => {
        console.log('handleInviteSelected')
    }

    const handleInviteAll = () => {
        console.log('handleInviteAll')
    }

    return(
        <Container>
            <Segment raised>
                <Header>{selected_exp.name}</Header>
                <p>{selected_exp.description}</p>
                <p><b># of rounds: </b><span>{selected_exp.config.rounds}</span></p>
                <p><b>resume configuration: </b><span>{selected_exp.config.resume}</span></p>
            </Segment>
            
            <Grid columns={2}>
                {/* column for displaying users who have yet to complete an experiment */}
                <Grid.Column>
                <Header as='h3'>Pending Invites</Header>
                    {selected_exp.pending.length > 0 ?
                        <List divided relaxed>
                        {selected_exp.pending.map(invite => {
                            return (<List.Item>{fake_data.user(invite.user).email}</List.Item>)
                        })}
                        </List>
                    :  
                        <p>There are no pending invites for this experiment!</p>
                    }
                </Grid.Column>
                
                {/* column for displaying users who have not been invited to the experiment */}
                <Grid.Column>
                    <Header as='h3'>Users available to invite</Header>
                    {selected_exp.usersNotInvited.length > 0 ?
                    <Container>
                        <Button content='Invite selected users' onClick={handleInviteSelected}/>
                        <Button content='Invite all users' onClick={handleInviteAll}/>
                        <List divided relaxed>
                        {selected_exp.usersNotInvited.map(user => {
                            return (<List.Item>
                                <Checkbox value={user} onChange={(e, data) => handleCheckbox(data)}/>
                                <List.Content floated='left'>{user.email}</List.Content>
                                <List.Content floated='right'>
                                    <Button content='Invite' onClick={handleInviteUser}/>
                                </List.Content>
                            </List.Item>)
                        })}
                        </List>
                    </Container>
                    :
                    <p>There are no users available to invite!</p>
                    }
                </Grid.Column>
            </Grid>
        </Container>
    )
}

export default ExperimentDetails;