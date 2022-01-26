import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import Grid from '@mui/material/Grid'
import HolderGroup from "../components/HolderGroup"
import StatusTab from "../components/StatusTab"
import ButtonsTab from "../components/ButtonsTab"
import TileGroup from "../components/TileGroup"
import Alert from '@mui/material/Alert';
import AlertTitle from '@mui/material/AlertTitle';
import { useState } from "react";
import styled from 'styled-components';

const UpperHolders = styled(Grid)`
    @media (max-width: 500px) {
        margin-top: 10px;
        margin-left: 10px;
        padding-top : 20px;
    }
`

export default function Calculator(){
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
    const [alert, setAlert] = useState(null)
    const [winningTile, setWinningTile] = useState(0)

    function padList() {
        let newTileList = tilesList
        let total = 14
        let count = tilesList.length
        for (let i = count; i < total; i++) {
            newTileList = ([...newTileList, 0])
        }
        return newTileList
    }

    return(
        <>
            <Navbar/>
            {
                alert ? <Alert severity="error" onClose={() => setAlert(null)} >{alert}</Alert> : <></>
            }

            {
                result.length !== 0 ? 
                <Alert severity="success">
                    <AlertTitle>You Have Won!</AlertTitle>
                    Total Score is {result.Score} <strong>[{result.Names}]</strong>
                </Alert> : <></>
            }

            <UpperHolders container sx={{pt: 10}}>
                <Grid item sm={3}>
                    <StatusTab openStatus1={openStatus1} setOpenStatus1={setOpenStatus1} openStatus2={openStatus2} setOpenStatus2={setOpenStatus2}
                        openStatus3={openStatus3} setOpenStatus3={setOpenStatus3} openStatus4={openStatus4} setOpenStatus4={setOpenStatus4}
                        openStatus5={openStatus5} setOpenStatus5={setOpenStatus5} kongStatus1={kongStatus1} setKongStatus1={setKongStatus1}
                        kongStatus2={kongStatus2} setKongStatus2={setKongStatus2} kongStatus3={kongStatus3} setKongStatus3={setKongStatus3}
                        kongStatus4={kongStatus4} setKongStatus4={setKongStatus4}
                    />
                </Grid>

                <Grid item sm={6}>
                    <HolderGroup tiles={padList(tilesList)} winningTile={winningTile} setWinningTile={setWinningTile}/>
                </Grid>
                
                <Grid item sm={3}>
                    <ButtonsTab tiles={tilesList} setTile={setTile} openStatus1={openStatus1} openStatus2={openStatus2} openStatus3={openStatus3}  
                        openStatus4={openStatus4} openStatus5={openStatus5}  kongStatus1={kongStatus1} kongStatus2={kongStatus2}  
                        kongStatus3={kongStatus3} kongStatus4={kongStatus4} alert={alert} setAlert={setAlert} setResult={setResult} winningTile={winningTile}
                    />
                </Grid>
            </UpperHolders>

            <Grid container direction="column" justifyContent="center" alignItems="center">
                <TileGroup setTile={setTile} tiles={tilesList}/>
            </Grid>

            <Footer />
        </>
    )
}