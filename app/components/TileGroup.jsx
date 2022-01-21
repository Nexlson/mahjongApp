import Tile from './Tile'
import range from './utils'

export default function TileGroup() {
    const tiles = range(1, 42);

    return(
        <>
            {tiles.map((tile, i) => (
                <Tile id={tile} key={i}/>
            ))}
        </>
    )
}
