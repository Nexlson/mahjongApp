import Button from '@mui/material/Button';
import styled from 'styled-components';

const StatusButtonsStyle = styled(Button)`
    margin: 15px 15px;
    
    @media (max-width: 500px) {
        min-width: 5px;
        margin: 5px 5px;
        font-size: 10px;
    }
`

export default function StatusButtons(props) {
    function onClickStatus(status, setStatus) {
        setStatus(!status)
    }

    return (
        <StatusButtonsStyle id={props.id} variant="contained" color={props.status ? "success" : "primary"} 
            onClick={()=> onClickStatus(props.status, props.setStatus)}> {props.name} </StatusButtonsStyle>
    )
}