import React from 'react';
import 'semantic-ui-css/semantic.min.css'

import * as fake_data from './apiFunctions.js';
import {Container, Header, Grid, Segment, List} from 'semantic-ui-react';

const ParticipantDetails = ({user_id}) => {

    let selected_user = fake_data.user(user_id);

    return(
        <Container>
            <Segment raised>
                <p><b>Email: </b><span>{selected_user.email}</span></p>
                <p><b>Source of registration: </b><span>{selected_user.source}</span></p>
            </Segment>
            
            <Grid columns={2}>
                {/* column for displaying experiments the participant has yet to complete */}
                <Grid.Column>
                <Header as='h3'>Experiments to complete</Header>
                {selected_user.invites.length > 0 ?
                <List divided relaxed>
                    {selected_user.invites.map(i => {
                        return (
                        <List.Item>
                            <List.Header>{fake_data.experiment(i.experiment).name}</List.Header>
                            <List.Description>{`Created at: `+i.created_at}</List.Description>
                        </List.Item>
                        )
                    })
                    }
                </List>
                :
                <p>This participant has no pending invites!</p>
                }
                </Grid.Column>
                
                {/* column for displaying the participant's results */}
                <Grid.Column>
                    <Header as='h3'>Experiment results</Header>
                    {selected_user.results.length > 0 ?
                    <List divided relaxed>
                        {selected_user.results.map(r => {
                            return (
                            <List.Item>
                                <List.Header>{fake_data.experiment(r.experiment).name}</List.Header>
                                <List.Description>{`Completed: `+r.created_at}</List.Description>
                            </List.Item>
                            )
                        })
                        }
                    </List>
                    :
                    <p>This participant has not completed any experiments!</p>
                    }
                </Grid.Column>
            </Grid>
        </Container>
    )
}

export default ParticipantDetails;