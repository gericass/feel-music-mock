package digimartapp.gericass.com.feelmucicmock.Track

import android.content.Context
import android.widget.TextView
import android.databinding.adapters.TextViewBindingAdapter.setText
import android.support.design.widget.CoordinatorLayout.Behavior.setTag
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.ImageView
import com.bumptech.glide.Glide
import digimartapp.gericass.com.feelmucicmock.R


/**
* Created by keita on 2018/06/11.
*/

class TrackCardAdapter(context: Context) : ArrayAdapter<Track>(context, 0) {

    override fun getView(position: Int, contentView: View?, parent: ViewGroup): View {
        var contentView = contentView
        val holder: ViewHolder

        if (contentView == null) {
            val inflater = LayoutInflater.from(context)
            contentView = inflater.inflate(R.layout.track_card, parent, false)
            holder = ViewHolder(contentView)
            contentView!!.tag = holder
        } else {
            holder = contentView.tag as ViewHolder
        }

        val spot = getItem(position)

        holder.name.text = spot.name
        holder.comment.text = spot.comment
        Glide.with(context).load(spot.url).into(holder.image)

        return contentView
    }

    private class ViewHolder(view: View) {
        var name: TextView = view.findViewById(R.id.track_name)
        var comment: TextView = view.findViewById(R.id.track_comment)
        var image: ImageView = view.findViewById(R.id.track_image) as ImageView

    }

}