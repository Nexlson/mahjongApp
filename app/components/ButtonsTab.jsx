import CameraAltIcon from '@mui/icons-material/CameraAlt';
import DeleteIcon from '@mui/icons-material/Delete';
import CalculateIcon from '@mui/icons-material/Calculate';
import RemoveIcon from '@mui/icons-material/Remove';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid'
import axios from 'axios'
import range from './utils'

export default function ButtonsTab(props) {
    const deleteTile = () => {
        props.setTile(props.tiles.slice(0, -1))
    }

    const clearTiles = () => {
        props.setTile([])
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

        // URL = "https://localhost:3500/api/v1/calculator"
            // let group1 = props.tiles.slice(0, 3)
            // let group2 = props.tiles.slice(3, 6)
            // let group3 = props.tiles.slice(6, 9)
            // let group4 = props.tiles.slice(9, 12)
            // let group5 = props.tiles.slice(12, 14)
        // let data = {
        //     "Grouped": [
        //         [group1, props.openStatus1, checkPong(group1), props.kongStatus1, checkChi(group1), false, checkPattern(group1)],
        //         [group2, props.openStatus2, checkPong(group2), props.kongStatus2, checkChi(group2), false, checkPattern(group2)],
        //         [group3, props.openStatus3, checkPong(group3), props.kongStatus3, checkChi(group3), false, checkPattern(group3)],
        //         [group4, props.openStatus4, checkPong(group4), props.kongStatus4, checkChi(group4), false, checkPattern(group4)],
        //         [group5, props.openStatus5, false, false, false, checkPong(group5), checkPattern(group5)]
        //     ],
        //     "Ungrouped": props.tiles,
        //     "Won": false
        // }
        // // to Json
        // data = JSON.parse(data)
        // // send to back end
        // axios({
        //     method: "POST",
        //     url: URL,
        //     data: {
        //         data
        //     }
        // })
        // .then(data=> console.log(data))
        // .catch(err=> console.log(err))

        console.log("Calculate Now")
    }

    return (
        <>
            <Grid container direction="column" justifyContent="center" alignItems="flex-start">
                <Button variant="contained" disabled endIcon={<CameraAltIcon />} sx={{mb: 5}}> 
                    Camera
                </Button>

                <Button variant="contained" endIcon={<DeleteIcon />} sx={{mb: 5}} onClick={() => deleteTile()}>
                    Delete
                </Button>


                <Button variant="contained" endIcon={<RemoveIcon />} sx={{mb: 5}} onClick={() => clearTiles()}>
                    Clear
                </Button>


                <Button variant="contained" endIcon={<CalculateIcon />} sx={{mb: 5}} onClick={() => calculateTiles()}>
                    Calculate
            </Button>
            </Grid>
        </>
    )
}