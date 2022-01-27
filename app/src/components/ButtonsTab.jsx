import CameraAltIcon from '@mui/icons-material/CameraAlt';
import DeleteIcon from '@mui/icons-material/Delete';
import CalculateIcon from '@mui/icons-material/Calculate';
import RemoveIcon from '@mui/icons-material/Remove';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid'
import axios from 'axios'
import { deleteTile, clearTiles, checkPong, checkPattern, checkChi } from '../utils/functions'
import { backendURL } from '../utils/defaults'
import styled from 'styled-components';

const ButtonTab = styled(Button)`
    margin: 20px 20px;

    span {
        margin-left: 0px;
    }

    @media (max-width: 500px) {
        width: 5px !important;
        height: 30px !important;
        margin-bottom: 10px;
        margin-left: 13px;
        justify-items: center;

      
        span {
            margin-left: 0px;
        }
    }
`

export default function ButtonsTab(props) {

    const calculateTiles = () => {
        if (props.tiles.length < 14){
            props.setAlert("Missing tiles, please try again!")
            return 
        }

        if (props.winningTile === 0) {
            props.setAlert("Missing winning tiles, please try again!")
            return
        }
    
        let group1 = props.tiles.slice(0, 3)
        let group2 = props.tiles.slice(3, 6)
        let group3 = props.tiles.slice(6, 9)
        let group4 = props.tiles.slice(9, 12)
        let group5 = props.tiles.slice(12, 14)
        let data = {
            "Grouped": [
                {"Tiles":group1, "Open":props.openStatus1, "Pong":checkPong(group1), "Kong":props.kongStatus1, "Chi":checkChi(group1), "Pair":false, "Pattern":checkPattern(group1)},
                {"Tiles":group2, "Open":props.openStatus2, "Pong":checkPong(group2), "Kong":props.kongStatus2, "Chi":checkChi(group2), "Pair":false, "Pattern":checkPattern(group2)},
                {"Tiles":group3, "Open":props.openStatus3, "Pong":checkPong(group3), "Kong":props.kongStatus3, "Chi":checkChi(group3), "Pair":false, "Pattern":checkPattern(group3)},
                {"Tiles":group4, "Open":props.openStatus4, "Pong":checkPong(group4), "Kong":props.kongStatus4, "Chi":checkChi(group4), "Pair":false, "Pattern":checkPattern(group4)},
                {"Tiles":group5, "Open":props.openStatus5, "Pong":false, "Kong":false, "Chi":false, "Pair":checkPong(group5), "Pattern":checkPattern(group5)}
            ],
            "Ungrouped": props.tiles,
            "Won": false,
            "WiningTile": props.winningTile - 1,
        }
        
        // send to back end
        axios({
            method: "POST",
            url: backendURL,
            data: data
        })
        .then(data=> {
            let result = data.data.result
            if (result.Score === 0) {
                props.setAlert("Incorrect inputs, please try again!")
                return
            }
            props.setResult(result)
        })
        .catch(err=> console.log(err))
    }

    return (
        <>
            <Grid container direction="column" justifyContent="center" alignItems="flex-start">
                <ButtonTab variant="contained" disabled endIcon={<CameraAltIcon />} />
                <ButtonTab variant="contained" endIcon={<DeleteIcon />} onClick={() => deleteTile(props.tiles, props.setTile)} />
                <ButtonTab variant="contained" endIcon={<RemoveIcon />} onClick={() => clearTiles(props.setTile, props.setResult, props.setOpenStatus1, props.setOpenStatus2, 
                    props.setOpenStatus3, props.setOpenStatus4, props.setOpenStatus5, props.setKongStatus1, props.setKongStatus2, props.setKongStatus3, props.setKongStatus4, props.setWinningTile)} />
                <ButtonTab variant="contained" endIcon={<CalculateIcon />} onClick={() => calculateTiles()} />
            </Grid>
        </>
    )
}