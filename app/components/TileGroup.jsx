import Tile from './Tile'
import range from './utils'
import Grid from '@mui/material/Grid'

export default function TileGroup(props) {
    const tilesRow1 = range(1, 9);
    const tilesRow2 = range(10, 18);
    const tilesRow3 = range(19, 27);
    const tilesRow4 = range(28, 34);
    const tilesRowList = [tilesRow1, tilesRow2, tilesRow3, tilesRow4]

    return(
        <>
            {
                tilesRowList.map((row, index) => (
                    <Grid container direction="row" justifyContent="center" alignItems="center" key={index}>
                    {row.map((tile, i) => (
                        <Grid item xs={1} key={i}> 
                            <Tile id={tile} key={i} setTiles={props.setTile} tiles={props.tiles}/>
                        </Grid>
                    ))}
                </Grid>
                ))
            }
        </>
    )
}
