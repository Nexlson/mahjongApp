import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import styled from 'styled-components';

const Status = styled(Grid)`
    margin-left: 15vw;
    padding-top: 10px;
    @media (max-width: 500px) {
        padding-top: 15px;
        margin-bottom: 20px;
        margin-left: 0;
    }
`

const StatusButton = styled(Button)`
    @media (max-width: 500px) {
        min-width: 5px;
        margin: 0 10px 5px !important;
        font-size: 10px;
    }
`

export default function StatusTab(props) {
    
    function onClickStatus(status, setStatus) {
        setStatus(!status)
    }
    return (
        <>  
            <Grid container direction="column" alignItems="stretch">
                <Status container direction="row" alignItems="center">
                    <StatusButton id={1} variant="contained" sx={{mr: 5, mb: 3, mt: 3}}  color={props.openStatus1 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus1, props.setOpenStatus1)}> 开 </StatusButton>
                    <StatusButton id={2} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus1 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus1, props.setKongStatus1)}> 碰 </StatusButton>
                </Status>
                <Status container direction="row" justifyContent="flex-start" alignItems="center">
                    <StatusButton id={3} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus2 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus2, props.setOpenStatus2)}> 开 </StatusButton>
                    <StatusButton id={4} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus2 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus2, props.setKongStatus2)}> 碰 </StatusButton>
                </Status>
                <Status container direction="row" justifyContent="flex-start" alignItems="center">
                    <StatusButton id={5} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus3 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus3, props.setOpenStatus3)}> 开 </StatusButton>
                    <StatusButton id={6} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus3 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus3, props.setKongStatus3)}> 碰 </StatusButton>
                </Status>
                <Status container direction="row" justifyContent="flex-start" alignItems="center">
                    <StatusButton id={7} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus4 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus4, props.setOpenStatus4)}> 开 </StatusButton>
                    <StatusButton id={8} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus4 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus4, props.setKongStatus4)}> 碰 </StatusButton>
                </Status>
                <Status container direction="row" justifyContent="flex-start" alignItems="center">
                    <StatusButton id={9} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus5 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus5, props.setOpenStatus5)}> 开 </StatusButton>
                </Status>
            </Grid>
        </>
    )
}