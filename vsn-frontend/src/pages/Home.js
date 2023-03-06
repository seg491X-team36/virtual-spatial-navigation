import React, {useState} from 'react';

import 'semantic-ui-css/semantic.min.css'
import { Button, ButtonGroup } from 'semantic-ui-react';

const Home = () => {
    return(
        <>
        <h1>Welcome!</h1>
        <h2>What would you like to do?</h2>
        <ButtonGroup>
            <Button>View Experiments</Button>
            <Button>View Participants</Button>
            <Button>Rejected Users</Button>
            <Button>View Sign Ups</Button>
        </ButtonGroup>
        </>
    )
};

export default Home;