import Holder from './Holder'
import Grid from '@mui/material/Grid'

export default function HolderGroup(props) {
    const holder1 = props.tiles.slice(0, 3)
    const holder2 = props.tiles.slice(3, 6)
    const holder3 = props.tiles.slice(6, 9)
    const holder4 = props.tiles.slice(9, 12)
    const holder5 = props.tiles.slice(12, 14)
    const holderList = [holder1, holder2, holder3, holder4, holder5]
    let count = 0
    
    return (
        <>  
            {
                holderList.map((holder, i) => (
                    <Grid container justifyContent="center" alignItems="center" key={i}>
                        {holder.map((hold, index) => (
                            <Holder key={index} idHold={count += 1} id={hold} image={props.tiles} winningTile={props.winningTile} setWinningTile={props.setWinningTile}/>
                            ))}
                    </Grid>
                ))
            }
        </>
    )
}