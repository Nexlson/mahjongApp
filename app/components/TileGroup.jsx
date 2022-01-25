import Tile from './Tile'
import range from './utils'

export default function TileGroup(props) {
    const tilesRow1 = range(1, 9);
    const tilesRow2 = range(10, 18);
    const tilesRow3 = range(19, 27);
    const tilesRow4 = range(28, 34);
    const tiles = range(1,34)
    const tilesRowList = [tilesRow1, tilesRow2, tilesRow3, tilesRow4]

    return(
        <>
            {
                tiles.map((tile, index) => (
                    <Tile id={tile} key={index} setTiles={props.setTile} tiles={props.tiles}/>
                ))
            }
        </>
    )
}
