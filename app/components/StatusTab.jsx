import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid'

export default function StatusTab(props) {
    
    function onClickStatus(status, setStatus) {
        setStatus(!status)
    }

    return (
        <>  
            <Grid container direction="column" alignItems="stretch">
                <Grid container direction="row" justifyContent="flex-end" alignItems="center">
                    <Button id={1} variant="contained" sx={{mr: 5, mb: 3, mt: 3}}  color={props.openStatus1 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.openStatus1, props.setOpenStatus1)}> Open </Button>
                    <Button id={2} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus1 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.kongStatus1, props.setKongStatus1)}> Kong </Button>
                </Grid>
                <Grid container direction="row" justifyContent="flex-end" alignItems="center">
                    <Button id={3} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus2 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.openStatus2, props.setOpenStatus2)}> Open </Button>
                    <Button id={4} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus2 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.kongStatus2, props.setKongStatus2)}> Kong </Button>
                </Grid>
                <Grid container direction="row" justifyContent="flex-end" alignItems="center">
                    <Button id={5} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus3 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.openStatus3, props.setOpenStatus3)}> Open </Button>
                    <Button id={6} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus3 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.kongStatus3, props.setKongStatus3)}> Kong </Button>
                </Grid>
                <Grid container direction="row" justifyContent="flex-end" alignItems="center">
                    <Button id={7} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus4 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.openStatus4, props.setOpenStatus4)}> Open </Button>
                    <Button id={8} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus4 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.kongStatus4, props.setKongStatus4)}> Kong </Button>
                </Grid>
                <Grid container direction="row" justifyContent="flex-end" alignItems="center">
                    <Button id={9} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus5 ? "secondary" : "primary"} onClick={()=> onClickStatus(props.openStatus5, props.setOpenStatus5)}> Open </Button>
                </Grid>
            </Grid>
        </>
    )
}