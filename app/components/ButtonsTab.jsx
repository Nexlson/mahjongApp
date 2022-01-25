import CameraAltIcon from '@mui/icons-material/CameraAlt';
import DeleteIcon from '@mui/icons-material/Delete';
import CalculateIcon from '@mui/icons-material/Calculate';
import RemoveIcon from '@mui/icons-material/Remove';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid'
import axios from 'axios'
import range from './utils'
import styles from '../styles/ButtonTab.module.css'

export default function ButtonsTab(props) {
    const deleteTile = () => {
        props.setTile(props.tiles.slice(0, -1))
    }

    const clearTiles = () => {
        props.setTile([])
        props.setResult([])
    }

    const checkPong = arr => arr.every( v => v === arr[0] )

    function checkChi(array) {
        var i = 2, d;
        while (i < array.length) {
            d = array[i - 1] - array[i - 2];
            if (Math.abs(d) === 1 && d === array[i] - array[i - 1]) {
                return true;
            }
            i++;
        }
        return false;
    }

    function onlyUnique(value, index, self) {
        return self.indexOf(value) === index;
      }

    function checkPattern(arr) {
        // remove duplicates in arr
        let removed = arr.filter(onlyUnique)

        // check which pattern class
        let pattern1 = range(1,9)
        let pattern2 = range(10,18)
        let pattern3 = range(19, 27)
        let pattern4 = range(28, 31)
        let pattern5 = range(32,34)

        if(removed.every(r => pattern1.includes(r))){
            return 1
        }else if(removed.every(r => pattern2.includes(r))){
            return 2
        }else if(removed.every(r => pattern3.includes(r))){
            return 3
        }else if(removed.every(r => pattern4.includes(r))){
            return 4
        }else if(removed.every(r => pattern5.includes(r))){
            return 5
        }else{
            return 6
        }
    }

    const calculateTiles = () => {
        if (props.tiles.length < 14){
            props.setAlert(true)
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
            "Won": false
        }

        // send to back end
        axios({
            method: "POST",
            url: "http://localhost/api/v1/calculator",
            data: data
        })
        .then(data=> {
            let result = data.data.result
            props.setResult(result)
        })
        .catch(err=> console.log(err))


    }

    return (
        <>
            <Grid container direction="column" justifyContent="center" alignItems="flex-start">
                <Button variant="contained" className={styles.button} disabled endIcon={<CameraAltIcon />}> 
                </Button>

                <Button variant="contained" className={styles.button} endIcon={<DeleteIcon />} onClick={() => deleteTile()}>
                </Button>


                <Button variant="contained" className={styles.button} endIcon={<RemoveIcon />} onClick={() => clearTiles()}>
                </Button>


                <Button variant="contained" className={styles.button} endIcon={<CalculateIcon />} onClick={() => calculateTiles()}>
                </Button>
            </Grid>
        </>
    )
}