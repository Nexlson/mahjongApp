// import Image from 'next/image'
import Image from 'material-ui-image'

export default function Tile(props) {
    const imageClick = (id) => {
        props.setTiles([...props.tiles, id])
    }

    return(
        <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.id} width={80} height={80} onClick={()=> imageClick(props.id)}/>
    )
}