import Image from 'next/image'

export default function Holder(props) {
    return(
        props.id == 0 ? (<div className="holder"><Image src={"/grey.jpg"} alt={"Tile"+props.id} width={80} height={80}/></div>) : 
        (<div className="holder"><Image src={"/tiles/"+props.id+".jpg"} alt={"Tile"+props.key} width={80} height={80}/></div>)
    )
}