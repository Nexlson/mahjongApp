import Image from 'next/image'

export default function Holder(props) {
    return(
        props.id == 0 ? <Image src={"/grey.jpg"} alt={"Tile"+props.id} width={80} height={80}/> : 
        <Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.key} width={80} height={80}/>
    )
}