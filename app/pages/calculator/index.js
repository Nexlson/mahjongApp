import TileGroup from "../../components/TileGroup"
import HolderGroup from "../../components/HolderGroup"
import Footer from '../../components/Footer'
import NavBar from '../../components/Navbar'
import Grid from '@mui/material/Grid'
import ButtonsTab from '../../components/ButtonsTab'
import StatusTab from '../../components/StatusTab'
import { useState, useEffect } from "react";
import Alert from '@mui/material/Alert';

export default function Calculator() {
    const [result, setResult] = useState([])
    const [tilesList, setTile] = useState([])
    const [openStatus1, setOpenStatus1] = useState(false)
    const [openStatus2, setOpenStatus2] = useState(false)
    const [openStatus3, setOpenStatus3] = useState(false)
    const [openStatus4, setOpenStatus4] = useState(false)
    const [openStatus5, setOpenStatus5] = useState(false)
    const [kongStatus1, setKongStatus1] = useState(false)
    const [kongStatus2, setKongStatus2] = useState(false)
    const [kongStatus3, setKongStatus3] = useState(false)
    const [kongStatus4, setKongStatus4] = useState(false)
    const [alert, setAlert] = useState(false)

    useEffect(() => {
        setTimeout(() => {
          setAlert(false);
        }, 3000);
      }, []);  

    function padList() {
        let newTileList = tilesList
        let total = 14
        let count = tilesList.length
        for (let i = count; i < total; i++) {
            newTileList = ([...newTileList, 0])
        }
        return newTileList
    }

    return (
        <>
            <NavBar />
            {
                alert ? <Alert severity="error">Missing tiles — check again!</Alert> : <></>
            }
            <Grid container sx={{pt: 10}}>
                <Grid item sm={3}>
                    <StatusTab openStatus1={openStatus1} setOpenStatus1={setOpenStatus1} openStatus2={openStatus2} setOpenStatus2={setOpenStatus2}
                        openStatus3={openStatus3} setOpenStatus3={setOpenStatus3} openStatus4={openStatus4} setOpenStatus4={setOpenStatus4}
                        openStatus5={openStatus5} setOpenStatus5={setOpenStatus5} kongStatus1={kongStatus1} setKongStatus1={setKongStatus1}
                        kongStatus2={kongStatus2} setKongStatus2={setKongStatus2} kongStatus3={kongStatus3} setKongStatus3={setKongStatus3}
                        kongStatus4={kongStatus4} setKongStatus4={setKongStatus4}
                    />
                </Grid>
                <Grid item sm={6}>
                    <HolderGroup tiles={padList(tilesList)}/>
                </Grid>
                <Grid item sm={3}>
                    <ButtonsTab tiles={tilesList} setTile={setTile} openStatus1={openStatus1} openStatus2={openStatus2} openStatus3={openStatus3}  
                        openStatus4={openStatus4} openStatus5={openStatus5}  kongStatus1={kongStatus1} kongStatus2={kongStatus2}  
                        kongStatus3={kongStatus3} kongStatus4={kongStatus4} alert={alert} setAlert={setAlert}
                    />
                </Grid>
            </Grid>
        
            <Grid container direction="row" justifyContent="center" alignItems="center" sx={{pt: 10, pb: 10}}>
                <Grid container item xs={6} spacing={2}>
                    <TileGroup setTile={setTile} tiles={tilesList}/>
                </Grid>
            </Grid>
            <Footer />
        </>
    )
}
