import Image from 'next/image'

export default function Tile(props) {
    return(
        <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.id} width={80} height={80} />
    )
}