package digimartapp.gericass.com.feelmucicmock

import android.content.Context
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.widget.LinearLayout
import digimartapp.gericass.com.feelmucicmock.constants.MAINFRAGMENT
import android.databinding.adapters.TextViewBindingAdapter.setText
import android.hardware.Sensor
import android.hardware.SensorManager
import android.hardware.Sensor.TYPE_TEMPERATURE
import android.hardware.Sensor.TYPE_STEP_DETECTOR
import android.hardware.Sensor.TYPE_STEP_COUNTER
import android.hardware.Sensor.TYPE_SIGNIFICANT_MOTION
import android.hardware.Sensor.TYPE_ROTATION_VECTOR
import android.hardware.Sensor.TYPE_RELATIVE_HUMIDITY
import android.hardware.Sensor.TYPE_PROXIMITY
import android.hardware.Sensor.TYPE_PRESSURE
import android.hardware.Sensor.TYPE_POSE_6DOF
import android.hardware.Sensor.TYPE_ORIENTATION
import android.hardware.Sensor.TYPE_MOTION_DETECT
import android.hardware.Sensor.TYPE_MAGNETIC_FIELD_UNCALIBRATED
import android.hardware.Sensor.TYPE_MAGNETIC_FIELD
import android.hardware.Sensor.TYPE_LINEAR_ACCELERATION
import android.hardware.Sensor.TYPE_LIGHT
import android.hardware.Sensor.TYPE_HEART_RATE
import android.hardware.Sensor.TYPE_HEART_BEAT
import android.hardware.Sensor.TYPE_GYROSCOPE_UNCALIBRATED
import android.hardware.Sensor.TYPE_GYROSCOPE
import android.hardware.Sensor.TYPE_GRAVITY
import android.hardware.Sensor.TYPE_GEOMAGNETIC_ROTATION_VECTOR
import android.hardware.Sensor.TYPE_GAME_ROTATION_VECTOR
import android.hardware.Sensor.TYPE_DEVICE_PRIVATE_BASE
import android.hardware.Sensor.TYPE_AMBIENT_TEMPERATURE
import android.hardware.Sensor.TYPE_ACCELEROMETER
import android.hardware.SensorEventListener
import android.widget.TextView
import android.support.v4.view.ViewCompat.setAlpha
import android.databinding.adapters.TextViewBindingAdapter.setText
import android.hardware.SensorEvent
import android.widget.Toast


class MainActivity : AppCompatActivity(), MainNavigator, SensorEventListener {

    private lateinit var mMainViewModel: MainViewModel

    private var sensorManager: SensorManager? = null

    private val sensorList = intArrayOf(Sensor.TYPE_ACCELEROMETER, Sensor.TYPE_AMBIENT_TEMPERATURE, Sensor.TYPE_DEVICE_PRIVATE_BASE, Sensor.TYPE_GAME_ROTATION_VECTOR, Sensor.TYPE_GEOMAGNETIC_ROTATION_VECTOR, Sensor.TYPE_GRAVITY, Sensor.TYPE_GYROSCOPE, Sensor.TYPE_GYROSCOPE_UNCALIBRATED, Sensor.TYPE_HEART_BEAT, Sensor.TYPE_HEART_RATE, Sensor.TYPE_LIGHT, Sensor.TYPE_LINEAR_ACCELERATION, Sensor.TYPE_MAGNETIC_FIELD, Sensor.TYPE_MAGNETIC_FIELD_UNCALIBRATED, Sensor.TYPE_MOTION_DETECT, Sensor.TYPE_ORIENTATION, Sensor.TYPE_POSE_6DOF, Sensor.TYPE_PRESSURE, Sensor.TYPE_PROXIMITY, Sensor.TYPE_RELATIVE_HUMIDITY, Sensor.TYPE_ROTATION_VECTOR, Sensor.TYPE_SIGNIFICANT_MOTION, Sensor.TYPE_STEP_COUNTER, Sensor.TYPE_STEP_DETECTOR, Sensor.TYPE_TEMPERATURE)

    private val sensorNameList = arrayOf("ACCELEROMETER", "AMBIENT_TEMPERATURE", "DEVICE_PRIVATE_BASE", "GAME_ROTATION_VECTOR", "GEOMAGNETIC_ROTATION_VECTOR", "GRAVITY", "GYROSCOPE", "GYROSCOPE_UNCALIBRATED", "HEART_BEAT", "HEART_RATE", "LIGHT", "LINEAR_ACCELERATION", "MAGNETIC_FIELD", "MAGNETIC_FIELD_UNCALIBRATED", "MOTION_DETECT", "ORIENTATION", "POSE_6DOF", "PRESSURE", "PROXIMITY", "RELATIVE_HUMIDITY", "ROTATION_VECTOR", "SIGNIFICANT_MOTION", "STEP_COUNTER", "STEP_DETECTOR", "TEMPERATURE")


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val mainFragment = MainFragment()
        mMainViewModel = MainViewModel(this, applicationContext)
        mainFragment.mViewmodel = mMainViewModel

        sensorManager = getSystemService(Context.SENSOR_SERVICE) as SensorManager

        val transaction = supportFragmentManager.beginTransaction()
        transaction.replace(R.id.main_fragment_container, mainFragment)
        transaction.commit()
    }

    override fun getSensor() {
        val sensors = sensorManager!!.getSensorList(Sensor.TYPE_ALL)
        val strListbuf = StringBuffer("Sensor List:\n\n")
        var count = 0

        for (sensor in sensors) {
            count++
            val str = String.format(
                    "%s: %s\n", (count + 1).toString(), sensor.name)
            strListbuf.append(str)
        }
        mMainViewModel.sensors = strListbuf

    }

    override fun onDestroy() {
        mMainViewModel.onActivityDestroyed()
        super.onDestroy()
    }


    override fun onResume() {
        mMainViewModel.onActivityResumed()
        super.onResume()
    }


    override fun getActiveSensor() {
        val strbuf = StringBuffer("Sensor List:\n\n")

        for (i in sensorList.indices) {
            val sensor = sensorManager?.getDefaultSensor(sensorList[i])

            if (sensor != null) {
                val strTmp = String.format("%s: %s: 使用可能\n",
                        (i + 1).toString(), sensorNameList[i])
                sensorManager?.registerListener(this, sensor, sensorList[i])
                strbuf.append(strTmp)
            } else {
                val strTmp = String.format("%s: %s: XXX 不可\n",
                        (i + 1).toString(), sensorNameList[i])
                strbuf.append(strTmp)
            }
        }
        mMainViewModel.sensors = strbuf
    }

    override fun onAccuracyChanged(p0: Sensor?, p1: Int) {}

    override fun onSensorChanged(event: SensorEvent) {
        when (event.sensor.type) {
            Sensor.TYPE_LIGHT -> {
                // 現在の明るさを取得
                val lightValue = event.values[0].toString()

                mMainViewModel.sensorValue = lightValue
            }
        }
    }
}
