import 'semantic-ui-css/semantic.min.css'

import NavBar from '../components/NavBar';

const ViewSignUps = () => {
    return (
        <>
            <h1>Sign Ups</h1>
            <table class="ui celled table">
                <thead>
                    <tr><th>Email</th>
                        <th></th>
                        <th></th>
                    </tr></thead>
                <tbody>
                    <tr>
                        <td data-label="Email">cdame453@uottawa.ca</td>
                        <td data-label="Accept">Accept</td>
                        <td data-label="Reject">Reject</td>
                    </tr>
                    <tr>
                        <td data-label="Email">ascfb435@uottawa.ca</td>
                        <td data-label="Accept">Accept</td>
                        <td data-label="Reject">Reject</td>
                    </tr>
                    <tr>
                        <td data-label="Email">ngdfg765@uottawa.ca</td>
                        <td data-label="Accept">Accept</td>
                        <td data-label="Reject">Reject</td>
                    </tr>
                </tbody>
            </table>
        </>
    )
};

export default ViewSignUps;