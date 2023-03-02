import {useState} from 'react';
import 'semantic-ui-css/semantic.min.css';
import {Grid, Header, Button, Container, Icon} from 'semantic-ui-react';

const DownloadGame = () => {

    const[downloaded, setDownloaded] = useState(false);

    const handleDownloadClick = () => {
        setDownloaded(true);
        console.log('download game...');
    }

    return (
        <Grid divided='vertically' columns={2}>
            <Grid.Column>
                {/* can use this later to place instructions */}
            </Grid.Column>
            <Grid.Column>
                <Container textAlign='center'>
                    <Header as='h1'>Welcome</Header>
                    <p>Please click the button below to download the Virtual Spatial Navigation game:</p>
                    <Button onClick={handleDownloadClick} disabled={downloaded} icon labelPosition='right'>
                        Download Game   
                        <Icon name='download'/>
                    </Button>
                </Container>
            </Grid.Column>
        </Grid>
    );
}

export default DownloadGame;