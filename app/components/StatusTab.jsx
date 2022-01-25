import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import styles from '../styles/StatusTab.module.css';

export default function StatusTab(props) {
    
    function onClickStatus(status, setStatus) {
        setStatus(!status)
    }

    return (
        <>  
            <Grid container direction="column" alignItems="stretch">
                <Grid container className={styles.statusTab} direction="row" alignItems="center">
                    <Button id={1} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}}  color={props.openStatus1 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus1, props.setOpenStatus1)}> 开 </Button>
                    <Button id={2} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus1 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus1, props.setKongStatus1)}> 碰 </Button>
                </Grid>
                <Grid container className={styles.statusTab} direction="row" justifyContent="flex-start" alignItems="center">
                    <Button id={3} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus2 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus2, props.setOpenStatus2)}> 开 </Button>
                    <Button id={4} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus2 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus2, props.setKongStatus2)}> 碰 </Button>
                </Grid>
                <Grid container className={styles.statusTab} direction="row" justifyContent="flex-start" alignItems="center">
                    <Button id={5} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus3 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus3, props.setOpenStatus3)}> 开 </Button>
                    <Button id={6} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus3 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus3, props.setKongStatus3)}> 碰 </Button>
                </Grid>
                <Grid container className={styles.statusTab} direction="row" justifyContent="flex-start" alignItems="center">
                    <Button id={7} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus4 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus4, props.setOpenStatus4)}> 开 </Button>
                    <Button id={8} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.kongStatus4 ? "success" : "primary"} onClick={()=> onClickStatus(props.kongStatus4, props.setKongStatus4)}> 碰 </Button>
                </Grid>
                <Grid container className={styles.statusTab} direction="row" justifyContent="flex-start" alignItems="center">
                    <Button id={9} className={styles.statusButton} variant="contained" sx={{mr: 5, mb: 3, mt: 3}} color={props.openStatus5 ? "success" : "primary"} onClick={()=> onClickStatus(props.openStatus5, props.setOpenStatus5)}> 开 </Button>
                </Grid>
            </Grid>
        </>
    )
}