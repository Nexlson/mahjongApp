import Image from 'next/image'
import styles from '../styles/Tile.module.css'

export default function Tile(props) {
    const imageClick = (id) => {
        props.setTiles([...props.tiles, id])
    }

    return(
        <div className={styles.tile}>
            <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.id} width={80} height={80} onClick={()=> imageClick(props.id)}/>
        </div>
    )
}