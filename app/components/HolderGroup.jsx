import Holder from './Holder'
import Grid from '@mui/material/Grid'

export default function HolderGroup(props) {
    const holder1 = props.tiles.slice(0, 3)
    const holder2 = props.tiles.slice(3, 6)
    const holder3 = props.tiles.slice(6, 9)
    const holder4 = props.tiles.slice(9, 12)
    const holder5 = props.tiles.slice(12, 14)
    const holderList = [holder1, holder2, holder3, holder4, holder5]
    

    return (
        <>  
            {
                holderList.map((holder, i) => (
                    <Grid container spacing={{ xs: 2, md: 3 }} direction="row" justifyContent="center" alignItems="center" key={i}>
                    {holder.map((hold, index) => (
                        <Grid item key={index}>
                            <Holder id={hold} image={props.tiles}/>
                        </Grid>
                    ))}
                </Grid>
                ))
            }
        </>
    )
}