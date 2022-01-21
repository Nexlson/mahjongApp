import Image from 'next/image'

export default function Holder(props) {
    return(
        <Image src={"/grey.jpg"} alt={"Tile"+props.id} width={80} height={80} />
    )
}