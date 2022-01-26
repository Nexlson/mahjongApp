import Grid from '@mui/material/Grid';
import styled from 'styled-components';
import StatusButtons from './StatusButtons'

const Status = styled(Grid)`
    margin-left: 100px;
    padding-top: 10px;
    @media (max-width: 500px) {
        padding-top: 15px;
        margin-bottom: 20px;
        margin-left: 0;
    }
`

export default function StatusTab(props) {
    return (
        <>  
            <Status container direction="row" alignItems="center">
                <StatusButtons id={1} status={props.openStatus1} setStatus={props.setOpenStatus1} name={"开"}/>
                <StatusButtons id={2} status={props.kongStatus1} setStatus={props.setKongStatus1} name={"扛"}/>
            </Status>
            <Status container direction="row" alignItems="center">
                <StatusButtons id={3} status={props.openStatus2} setStatus={props.setOpenStatus2} name={"开"}/>
                <StatusButtons id={4} status={props.kongStatus2} setStatus={props.setKongStatus2} name={"扛"}/>
            </Status>
            <Status container direction="row" alignItems="center">
                <StatusButtons id={5} status={props.openStatus3} setStatus={props.setOpenStatus3} name={"开"}/>
                <StatusButtons id={6} status={props.kongStatus3} setStatus={props.setKongStatus3} name={"扛"}/>
            </Status>
            <Status container direction="row" alignItems="center">
                <StatusButtons id={7} status={props.openStatus4} setStatus={props.setOpenStatus4} name={"开"}/>
                <StatusButtons id={8} status={props.kongStatus4} setStatus={props.setKongStatus4} name={"扛"}/>
            </Status>
            <Status container direction="row" alignItems="center">
                <StatusButtons id={9} status={props.openStatus5} setStatus={props.setOpenStatus5} name={"开"}/>
            </Status>
        </>
    )
}