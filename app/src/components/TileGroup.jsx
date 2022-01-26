import Tile from './Tile'
import { range } from '../utils/functions'
import styled from 'styled-components';
import Grid from '@mui/material/Grid'

const TileHolders = styled(Grid)`
    margin-top: 20px;
    display: -webkit-box;
`

export default function TileGroup(props) {
    const tiles1 = range(1, 9)
    const tiles2 = range(10, 18)
    const tiles3 = range(19, 27)
    const tiles4 = range(28, 34)
    const tiles = [tiles1, tiles2, tiles3, tiles4]

    return(
        <>
            {
                tiles.map((tileGroup, index) => (
                    <TileHolders key={index} container direction="column" justifyContent="center" alignItems="center">
                        {
                            tileGroup.map((tile, index) => (
                                <Tile id={tile} key={index} setTiles={props.setTile} tiles={props.tiles}/>
                            ))
                        }
                    </TileHolders>
                ))
            }
        </>
    )
}
