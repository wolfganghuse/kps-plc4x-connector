//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type KnxDatapointType uint32

type IKnxDatapointType interface {
    Number() uint16
    Name() string
    DatapointMainType() KnxDatapointMainType
    Serialize(io utils.WriteBuffer) error
}

const(
    KnxDatapointType_DPT_UNKNOWN KnxDatapointType = 0
    KnxDatapointType_BOOL KnxDatapointType = 1
    KnxDatapointType_BYTE KnxDatapointType = 2
    KnxDatapointType_WORD KnxDatapointType = 3
    KnxDatapointType_DWORD KnxDatapointType = 4
    KnxDatapointType_LWORD KnxDatapointType = 5
    KnxDatapointType_USINT KnxDatapointType = 6
    KnxDatapointType_SINT KnxDatapointType = 7
    KnxDatapointType_UINT KnxDatapointType = 8
    KnxDatapointType_INT KnxDatapointType = 9
    KnxDatapointType_UDINT KnxDatapointType = 10
    KnxDatapointType_DINT KnxDatapointType = 11
    KnxDatapointType_ULINT KnxDatapointType = 12
    KnxDatapointType_LINT KnxDatapointType = 13
    KnxDatapointType_REAL KnxDatapointType = 14
    KnxDatapointType_LREAL KnxDatapointType = 15
    KnxDatapointType_CHAR KnxDatapointType = 16
    KnxDatapointType_WCHAR KnxDatapointType = 17
    KnxDatapointType_STRING KnxDatapointType = 18
    KnxDatapointType_WSTRING KnxDatapointType = 19
    KnxDatapointType_TIME KnxDatapointType = 20
    KnxDatapointType_LTIME KnxDatapointType = 21
    KnxDatapointType_DATE KnxDatapointType = 22
    KnxDatapointType_TIME_OF_DAY KnxDatapointType = 23
    KnxDatapointType_TOD KnxDatapointType = 24
    KnxDatapointType_DATE_AND_TIME KnxDatapointType = 25
    KnxDatapointType_DT KnxDatapointType = 26
    KnxDatapointType_DPT_Switch KnxDatapointType = 27
    KnxDatapointType_DPT_Bool KnxDatapointType = 28
    KnxDatapointType_DPT_Enable KnxDatapointType = 29
    KnxDatapointType_DPT_Ramp KnxDatapointType = 30
    KnxDatapointType_DPT_Alarm KnxDatapointType = 31
    KnxDatapointType_DPT_BinaryValue KnxDatapointType = 32
    KnxDatapointType_DPT_Step KnxDatapointType = 33
    KnxDatapointType_DPT_UpDown KnxDatapointType = 34
    KnxDatapointType_DPT_OpenClose KnxDatapointType = 35
    KnxDatapointType_DPT_Start KnxDatapointType = 36
    KnxDatapointType_DPT_State KnxDatapointType = 37
    KnxDatapointType_DPT_Invert KnxDatapointType = 38
    KnxDatapointType_DPT_DimSendStyle KnxDatapointType = 39
    KnxDatapointType_DPT_InputSource KnxDatapointType = 40
    KnxDatapointType_DPT_Reset KnxDatapointType = 41
    KnxDatapointType_DPT_Ack KnxDatapointType = 42
    KnxDatapointType_DPT_Trigger KnxDatapointType = 43
    KnxDatapointType_DPT_Occupancy KnxDatapointType = 44
    KnxDatapointType_DPT_Window_Door KnxDatapointType = 45
    KnxDatapointType_DPT_LogicalFunction KnxDatapointType = 46
    KnxDatapointType_DPT_Scene_AB KnxDatapointType = 47
    KnxDatapointType_DPT_ShutterBlinds_Mode KnxDatapointType = 48
    KnxDatapointType_DPT_DayNight KnxDatapointType = 49
    KnxDatapointType_DPT_Heat_Cool KnxDatapointType = 50
    KnxDatapointType_DPT_Switch_Control KnxDatapointType = 51
    KnxDatapointType_DPT_Bool_Control KnxDatapointType = 52
    KnxDatapointType_DPT_Enable_Control KnxDatapointType = 53
    KnxDatapointType_DPT_Ramp_Control KnxDatapointType = 54
    KnxDatapointType_DPT_Alarm_Control KnxDatapointType = 55
    KnxDatapointType_DPT_BinaryValue_Control KnxDatapointType = 56
    KnxDatapointType_DPT_Step_Control KnxDatapointType = 57
    KnxDatapointType_DPT_Direction1_Control KnxDatapointType = 58
    KnxDatapointType_DPT_Direction2_Control KnxDatapointType = 59
    KnxDatapointType_DPT_Start_Control KnxDatapointType = 60
    KnxDatapointType_DPT_State_Control KnxDatapointType = 61
    KnxDatapointType_DPT_Invert_Control KnxDatapointType = 62
    KnxDatapointType_DPT_Control_Dimming KnxDatapointType = 63
    KnxDatapointType_DPT_Control_Blinds KnxDatapointType = 64
    KnxDatapointType_DPT_Char_ASCII KnxDatapointType = 65
    KnxDatapointType_DPT_Char_8859_1 KnxDatapointType = 66
    KnxDatapointType_DPT_Scaling KnxDatapointType = 67
    KnxDatapointType_DPT_Angle KnxDatapointType = 68
    KnxDatapointType_DPT_Percent_U8 KnxDatapointType = 69
    KnxDatapointType_DPT_DecimalFactor KnxDatapointType = 70
    KnxDatapointType_DPT_Tariff KnxDatapointType = 71
    KnxDatapointType_DPT_Value_1_Ucount KnxDatapointType = 72
    KnxDatapointType_DPT_FanStage KnxDatapointType = 73
    KnxDatapointType_DPT_Percent_V8 KnxDatapointType = 74
    KnxDatapointType_DPT_Value_1_Count KnxDatapointType = 75
    KnxDatapointType_DPT_Status_Mode3 KnxDatapointType = 76
    KnxDatapointType_DPT_Value_2_Ucount KnxDatapointType = 77
    KnxDatapointType_DPT_TimePeriodMsec KnxDatapointType = 78
    KnxDatapointType_DPT_TimePeriod10Msec KnxDatapointType = 79
    KnxDatapointType_DPT_TimePeriod100Msec KnxDatapointType = 80
    KnxDatapointType_DPT_TimePeriodSec KnxDatapointType = 81
    KnxDatapointType_DPT_TimePeriodMin KnxDatapointType = 82
    KnxDatapointType_DPT_TimePeriodHrs KnxDatapointType = 83
    KnxDatapointType_DPT_PropDataType KnxDatapointType = 84
    KnxDatapointType_DPT_Length_mm KnxDatapointType = 85
    KnxDatapointType_DPT_UElCurrentmA KnxDatapointType = 86
    KnxDatapointType_DPT_Brightness KnxDatapointType = 87
    KnxDatapointType_DPT_Absolute_Colour_Temperature KnxDatapointType = 88
    KnxDatapointType_DPT_Value_2_Count KnxDatapointType = 89
    KnxDatapointType_DPT_DeltaTimeMsec KnxDatapointType = 90
    KnxDatapointType_DPT_DeltaTime10Msec KnxDatapointType = 91
    KnxDatapointType_DPT_DeltaTime100Msec KnxDatapointType = 92
    KnxDatapointType_DPT_DeltaTimeSec KnxDatapointType = 93
    KnxDatapointType_DPT_DeltaTimeMin KnxDatapointType = 94
    KnxDatapointType_DPT_DeltaTimeHrs KnxDatapointType = 95
    KnxDatapointType_DPT_Percent_V16 KnxDatapointType = 96
    KnxDatapointType_DPT_Rotation_Angle KnxDatapointType = 97
    KnxDatapointType_DPT_Length_m KnxDatapointType = 98
    KnxDatapointType_DPT_Value_Temp KnxDatapointType = 99
    KnxDatapointType_DPT_Value_Tempd KnxDatapointType = 100
    KnxDatapointType_DPT_Value_Tempa KnxDatapointType = 101
    KnxDatapointType_DPT_Value_Lux KnxDatapointType = 102
    KnxDatapointType_DPT_Value_Wsp KnxDatapointType = 103
    KnxDatapointType_DPT_Value_Pres KnxDatapointType = 104
    KnxDatapointType_DPT_Value_Humidity KnxDatapointType = 105
    KnxDatapointType_DPT_Value_AirQuality KnxDatapointType = 106
    KnxDatapointType_DPT_Value_AirFlow KnxDatapointType = 107
    KnxDatapointType_DPT_Value_Time1 KnxDatapointType = 108
    KnxDatapointType_DPT_Value_Time2 KnxDatapointType = 109
    KnxDatapointType_DPT_Value_Volt KnxDatapointType = 110
    KnxDatapointType_DPT_Value_Curr KnxDatapointType = 111
    KnxDatapointType_DPT_PowerDensity KnxDatapointType = 112
    KnxDatapointType_DPT_KelvinPerPercent KnxDatapointType = 113
    KnxDatapointType_DPT_Power KnxDatapointType = 114
    KnxDatapointType_DPT_Value_Volume_Flow KnxDatapointType = 115
    KnxDatapointType_DPT_Rain_Amount KnxDatapointType = 116
    KnxDatapointType_DPT_Value_Temp_F KnxDatapointType = 117
    KnxDatapointType_DPT_Value_Wsp_kmh KnxDatapointType = 118
    KnxDatapointType_DPT_Value_Absolute_Humidity KnxDatapointType = 119
    KnxDatapointType_DPT_Concentration_ygm3 KnxDatapointType = 120
    KnxDatapointType_DPT_TimeOfDay KnxDatapointType = 121
    KnxDatapointType_DPT_Date KnxDatapointType = 122
    KnxDatapointType_DPT_Value_4_Ucount KnxDatapointType = 123
    KnxDatapointType_DPT_LongTimePeriod_Sec KnxDatapointType = 124
    KnxDatapointType_DPT_LongTimePeriod_Min KnxDatapointType = 125
    KnxDatapointType_DPT_LongTimePeriod_Hrs KnxDatapointType = 126
    KnxDatapointType_DPT_VolumeLiquid_Litre KnxDatapointType = 127
    KnxDatapointType_DPT_Volume_m_3 KnxDatapointType = 128
    KnxDatapointType_DPT_Value_4_Count KnxDatapointType = 129
    KnxDatapointType_DPT_FlowRate_m3h KnxDatapointType = 130
    KnxDatapointType_DPT_ActiveEnergy KnxDatapointType = 131
    KnxDatapointType_DPT_ApparantEnergy KnxDatapointType = 132
    KnxDatapointType_DPT_ReactiveEnergy KnxDatapointType = 133
    KnxDatapointType_DPT_ActiveEnergy_kWh KnxDatapointType = 134
    KnxDatapointType_DPT_ApparantEnergy_kVAh KnxDatapointType = 135
    KnxDatapointType_DPT_ReactiveEnergy_kVARh KnxDatapointType = 136
    KnxDatapointType_DPT_ActiveEnergy_MWh KnxDatapointType = 137
    KnxDatapointType_DPT_LongDeltaTimeSec KnxDatapointType = 138
    KnxDatapointType_DPT_DeltaVolumeLiquid_Litre KnxDatapointType = 139
    KnxDatapointType_DPT_DeltaVolume_m_3 KnxDatapointType = 140
    KnxDatapointType_DPT_Value_Acceleration KnxDatapointType = 141
    KnxDatapointType_DPT_Value_Acceleration_Angular KnxDatapointType = 142
    KnxDatapointType_DPT_Value_Activation_Energy KnxDatapointType = 143
    KnxDatapointType_DPT_Value_Activity KnxDatapointType = 144
    KnxDatapointType_DPT_Value_Mol KnxDatapointType = 145
    KnxDatapointType_DPT_Value_Amplitude KnxDatapointType = 146
    KnxDatapointType_DPT_Value_AngleRad KnxDatapointType = 147
    KnxDatapointType_DPT_Value_AngleDeg KnxDatapointType = 148
    KnxDatapointType_DPT_Value_Angular_Momentum KnxDatapointType = 149
    KnxDatapointType_DPT_Value_Angular_Velocity KnxDatapointType = 150
    KnxDatapointType_DPT_Value_Area KnxDatapointType = 151
    KnxDatapointType_DPT_Value_Capacitance KnxDatapointType = 152
    KnxDatapointType_DPT_Value_Charge_DensitySurface KnxDatapointType = 153
    KnxDatapointType_DPT_Value_Charge_DensityVolume KnxDatapointType = 154
    KnxDatapointType_DPT_Value_Compressibility KnxDatapointType = 155
    KnxDatapointType_DPT_Value_Conductance KnxDatapointType = 156
    KnxDatapointType_DPT_Value_Electrical_Conductivity KnxDatapointType = 157
    KnxDatapointType_DPT_Value_Density KnxDatapointType = 158
    KnxDatapointType_DPT_Value_Electric_Charge KnxDatapointType = 159
    KnxDatapointType_DPT_Value_Electric_Current KnxDatapointType = 160
    KnxDatapointType_DPT_Value_Electric_CurrentDensity KnxDatapointType = 161
    KnxDatapointType_DPT_Value_Electric_DipoleMoment KnxDatapointType = 162
    KnxDatapointType_DPT_Value_Electric_Displacement KnxDatapointType = 163
    KnxDatapointType_DPT_Value_Electric_FieldStrength KnxDatapointType = 164
    KnxDatapointType_DPT_Value_Electric_Flux KnxDatapointType = 165
    KnxDatapointType_DPT_Value_Electric_FluxDensity KnxDatapointType = 166
    KnxDatapointType_DPT_Value_Electric_Polarization KnxDatapointType = 167
    KnxDatapointType_DPT_Value_Electric_Potential KnxDatapointType = 168
    KnxDatapointType_DPT_Value_Electric_PotentialDifference KnxDatapointType = 169
    KnxDatapointType_DPT_Value_ElectromagneticMoment KnxDatapointType = 170
    KnxDatapointType_DPT_Value_Electromotive_Force KnxDatapointType = 171
    KnxDatapointType_DPT_Value_Energy KnxDatapointType = 172
    KnxDatapointType_DPT_Value_Force KnxDatapointType = 173
    KnxDatapointType_DPT_Value_Frequency KnxDatapointType = 174
    KnxDatapointType_DPT_Value_Angular_Frequency KnxDatapointType = 175
    KnxDatapointType_DPT_Value_Heat_Capacity KnxDatapointType = 176
    KnxDatapointType_DPT_Value_Heat_FlowRate KnxDatapointType = 177
    KnxDatapointType_DPT_Value_Heat_Quantity KnxDatapointType = 178
    KnxDatapointType_DPT_Value_Impedance KnxDatapointType = 179
    KnxDatapointType_DPT_Value_Length KnxDatapointType = 180
    KnxDatapointType_DPT_Value_Light_Quantity KnxDatapointType = 181
    KnxDatapointType_DPT_Value_Luminance KnxDatapointType = 182
    KnxDatapointType_DPT_Value_Luminous_Flux KnxDatapointType = 183
    KnxDatapointType_DPT_Value_Luminous_Intensity KnxDatapointType = 184
    KnxDatapointType_DPT_Value_Magnetic_FieldStrength KnxDatapointType = 185
    KnxDatapointType_DPT_Value_Magnetic_Flux KnxDatapointType = 186
    KnxDatapointType_DPT_Value_Magnetic_FluxDensity KnxDatapointType = 187
    KnxDatapointType_DPT_Value_Magnetic_Moment KnxDatapointType = 188
    KnxDatapointType_DPT_Value_Magnetic_Polarization KnxDatapointType = 189
    KnxDatapointType_DPT_Value_Magnetization KnxDatapointType = 190
    KnxDatapointType_DPT_Value_MagnetomotiveForce KnxDatapointType = 191
    KnxDatapointType_DPT_Value_Mass KnxDatapointType = 192
    KnxDatapointType_DPT_Value_MassFlux KnxDatapointType = 193
    KnxDatapointType_DPT_Value_Momentum KnxDatapointType = 194
    KnxDatapointType_DPT_Value_Phase_AngleRad KnxDatapointType = 195
    KnxDatapointType_DPT_Value_Phase_AngleDeg KnxDatapointType = 196
    KnxDatapointType_DPT_Value_Power KnxDatapointType = 197
    KnxDatapointType_DPT_Value_Power_Factor KnxDatapointType = 198
    KnxDatapointType_DPT_Value_Pressure KnxDatapointType = 199
    KnxDatapointType_DPT_Value_Reactance KnxDatapointType = 200
    KnxDatapointType_DPT_Value_Resistance KnxDatapointType = 201
    KnxDatapointType_DPT_Value_Resistivity KnxDatapointType = 202
    KnxDatapointType_DPT_Value_SelfInductance KnxDatapointType = 203
    KnxDatapointType_DPT_Value_SolidAngle KnxDatapointType = 204
    KnxDatapointType_DPT_Value_Sound_Intensity KnxDatapointType = 205
    KnxDatapointType_DPT_Value_Speed KnxDatapointType = 206
    KnxDatapointType_DPT_Value_Stress KnxDatapointType = 207
    KnxDatapointType_DPT_Value_Surface_Tension KnxDatapointType = 208
    KnxDatapointType_DPT_Value_Common_Temperature KnxDatapointType = 209
    KnxDatapointType_DPT_Value_Absolute_Temperature KnxDatapointType = 210
    KnxDatapointType_DPT_Value_TemperatureDifference KnxDatapointType = 211
    KnxDatapointType_DPT_Value_Thermal_Capacity KnxDatapointType = 212
    KnxDatapointType_DPT_Value_Thermal_Conductivity KnxDatapointType = 213
    KnxDatapointType_DPT_Value_ThermoelectricPower KnxDatapointType = 214
    KnxDatapointType_DPT_Value_Time KnxDatapointType = 215
    KnxDatapointType_DPT_Value_Torque KnxDatapointType = 216
    KnxDatapointType_DPT_Value_Volume KnxDatapointType = 217
    KnxDatapointType_DPT_Value_Volume_Flux KnxDatapointType = 218
    KnxDatapointType_DPT_Value_Weight KnxDatapointType = 219
    KnxDatapointType_DPT_Value_Work KnxDatapointType = 220
    KnxDatapointType_DPT_Volume_Flux_Meter KnxDatapointType = 221
    KnxDatapointType_DPT_Volume_Flux_ls KnxDatapointType = 222
    KnxDatapointType_DPT_Access_Data KnxDatapointType = 223
    KnxDatapointType_DPT_String_ASCII KnxDatapointType = 224
    KnxDatapointType_DPT_String_8859_1 KnxDatapointType = 225
    KnxDatapointType_DPT_SceneNumber KnxDatapointType = 226
    KnxDatapointType_DPT_SceneControl KnxDatapointType = 227
    KnxDatapointType_DPT_DateTime KnxDatapointType = 228
    KnxDatapointType_DPT_SCLOMode KnxDatapointType = 229
    KnxDatapointType_DPT_BuildingMode KnxDatapointType = 230
    KnxDatapointType_DPT_OccMode KnxDatapointType = 231
    KnxDatapointType_DPT_Priority KnxDatapointType = 232
    KnxDatapointType_DPT_LightApplicationMode KnxDatapointType = 233
    KnxDatapointType_DPT_ApplicationArea KnxDatapointType = 234
    KnxDatapointType_DPT_AlarmClassType KnxDatapointType = 235
    KnxDatapointType_DPT_PSUMode KnxDatapointType = 236
    KnxDatapointType_DPT_ErrorClass_System KnxDatapointType = 237
    KnxDatapointType_DPT_ErrorClass_HVAC KnxDatapointType = 238
    KnxDatapointType_DPT_Time_Delay KnxDatapointType = 239
    KnxDatapointType_DPT_Beaufort_Wind_Force_Scale KnxDatapointType = 240
    KnxDatapointType_DPT_SensorSelect KnxDatapointType = 241
    KnxDatapointType_DPT_ActuatorConnectType KnxDatapointType = 242
    KnxDatapointType_DPT_Cloud_Cover KnxDatapointType = 243
    KnxDatapointType_DPT_PowerReturnMode KnxDatapointType = 244
    KnxDatapointType_DPT_FuelType KnxDatapointType = 245
    KnxDatapointType_DPT_BurnerType KnxDatapointType = 246
    KnxDatapointType_DPT_HVACMode KnxDatapointType = 247
    KnxDatapointType_DPT_DHWMode KnxDatapointType = 248
    KnxDatapointType_DPT_LoadPriority KnxDatapointType = 249
    KnxDatapointType_DPT_HVACContrMode KnxDatapointType = 250
    KnxDatapointType_DPT_HVACEmergMode KnxDatapointType = 251
    KnxDatapointType_DPT_ChangeoverMode KnxDatapointType = 252
    KnxDatapointType_DPT_ValveMode KnxDatapointType = 253
    KnxDatapointType_DPT_DamperMode KnxDatapointType = 254
    KnxDatapointType_DPT_HeaterMode KnxDatapointType = 255
    KnxDatapointType_DPT_FanMode KnxDatapointType = 256
    KnxDatapointType_DPT_MasterSlaveMode KnxDatapointType = 257
    KnxDatapointType_DPT_StatusRoomSetp KnxDatapointType = 258
    KnxDatapointType_DPT_Metering_DeviceType KnxDatapointType = 259
    KnxDatapointType_DPT_HumDehumMode KnxDatapointType = 260
    KnxDatapointType_DPT_EnableHCStage KnxDatapointType = 261
    KnxDatapointType_DPT_ADAType KnxDatapointType = 262
    KnxDatapointType_DPT_BackupMode KnxDatapointType = 263
    KnxDatapointType_DPT_StartSynchronization KnxDatapointType = 264
    KnxDatapointType_DPT_Behaviour_Lock_Unlock KnxDatapointType = 265
    KnxDatapointType_DPT_Behaviour_Bus_Power_Up_Down KnxDatapointType = 266
    KnxDatapointType_DPT_DALI_Fade_Time KnxDatapointType = 267
    KnxDatapointType_DPT_BlinkingMode KnxDatapointType = 268
    KnxDatapointType_DPT_LightControlMode KnxDatapointType = 269
    KnxDatapointType_DPT_SwitchPBModel KnxDatapointType = 270
    KnxDatapointType_DPT_PBAction KnxDatapointType = 271
    KnxDatapointType_DPT_DimmPBModel KnxDatapointType = 272
    KnxDatapointType_DPT_SwitchOnMode KnxDatapointType = 273
    KnxDatapointType_DPT_LoadTypeSet KnxDatapointType = 274
    KnxDatapointType_DPT_LoadTypeDetected KnxDatapointType = 275
    KnxDatapointType_DPT_Converter_Test_Control KnxDatapointType = 276
    KnxDatapointType_DPT_SABExcept_Behaviour KnxDatapointType = 277
    KnxDatapointType_DPT_SABBehaviour_Lock_Unlock KnxDatapointType = 278
    KnxDatapointType_DPT_SSSBMode KnxDatapointType = 279
    KnxDatapointType_DPT_BlindsControlMode KnxDatapointType = 280
    KnxDatapointType_DPT_CommMode KnxDatapointType = 281
    KnxDatapointType_DPT_AddInfoTypes KnxDatapointType = 282
    KnxDatapointType_DPT_RF_ModeSelect KnxDatapointType = 283
    KnxDatapointType_DPT_RF_FilterSelect KnxDatapointType = 284
    KnxDatapointType_DPT_StatusGen KnxDatapointType = 285
    KnxDatapointType_DPT_Device_Control KnxDatapointType = 286
    KnxDatapointType_DPT_ForceSign KnxDatapointType = 287
    KnxDatapointType_DPT_ForceSignCool KnxDatapointType = 288
    KnxDatapointType_DPT_StatusRHC KnxDatapointType = 289
    KnxDatapointType_DPT_StatusSDHWC KnxDatapointType = 290
    KnxDatapointType_DPT_FuelTypeSet KnxDatapointType = 291
    KnxDatapointType_DPT_StatusRCC KnxDatapointType = 292
    KnxDatapointType_DPT_StatusAHU KnxDatapointType = 293
    KnxDatapointType_DPT_CombinedStatus_RTSM KnxDatapointType = 294
    KnxDatapointType_DPT_LightActuatorErrorInfo KnxDatapointType = 295
    KnxDatapointType_DPT_RF_ModeInfo KnxDatapointType = 296
    KnxDatapointType_DPT_RF_FilterInfo KnxDatapointType = 297
    KnxDatapointType_DPT_Channel_Activation_8 KnxDatapointType = 298
    KnxDatapointType_DPT_StatusDHWC KnxDatapointType = 299
    KnxDatapointType_DPT_StatusRHCC KnxDatapointType = 300
    KnxDatapointType_DPT_CombinedStatus_HVA KnxDatapointType = 301
    KnxDatapointType_DPT_CombinedStatus_RTC KnxDatapointType = 302
    KnxDatapointType_DPT_Media KnxDatapointType = 303
    KnxDatapointType_DPT_Channel_Activation_16 KnxDatapointType = 304
    KnxDatapointType_DPT_OnOffAction KnxDatapointType = 305
    KnxDatapointType_DPT_Alarm_Reaction KnxDatapointType = 306
    KnxDatapointType_DPT_UpDown_Action KnxDatapointType = 307
    KnxDatapointType_DPT_HVAC_PB_Action KnxDatapointType = 308
    KnxDatapointType_DPT_DoubleNibble KnxDatapointType = 309
    KnxDatapointType_DPT_SceneInfo KnxDatapointType = 310
    KnxDatapointType_DPT_CombinedInfoOnOff KnxDatapointType = 311
    KnxDatapointType_DPT_ActiveEnergy_V64 KnxDatapointType = 312
    KnxDatapointType_DPT_ApparantEnergy_V64 KnxDatapointType = 313
    KnxDatapointType_DPT_ReactiveEnergy_V64 KnxDatapointType = 314
    KnxDatapointType_DPT_Channel_Activation_24 KnxDatapointType = 315
    KnxDatapointType_DPT_HVACModeNext KnxDatapointType = 316
    KnxDatapointType_DPT_DHWModeNext KnxDatapointType = 317
    KnxDatapointType_DPT_OccModeNext KnxDatapointType = 318
    KnxDatapointType_DPT_BuildingModeNext KnxDatapointType = 319
    KnxDatapointType_DPT_Version KnxDatapointType = 320
    KnxDatapointType_DPT_AlarmInfo KnxDatapointType = 321
    KnxDatapointType_DPT_TempRoomSetpSetF16_3 KnxDatapointType = 322
    KnxDatapointType_DPT_TempRoomSetpSetShiftF16_3 KnxDatapointType = 323
    KnxDatapointType_DPT_Scaling_Speed KnxDatapointType = 324
    KnxDatapointType_DPT_Scaling_Step_Time KnxDatapointType = 325
    KnxDatapointType_DPT_MeteringValue KnxDatapointType = 326
    KnxDatapointType_DPT_MBus_Address KnxDatapointType = 327
    KnxDatapointType_DPT_Colour_RGB KnxDatapointType = 328
    KnxDatapointType_DPT_LanguageCodeAlpha2_ASCII KnxDatapointType = 329
    KnxDatapointType_DPT_Tariff_ActiveEnergy KnxDatapointType = 330
    KnxDatapointType_DPT_Prioritised_Mode_Control KnxDatapointType = 331
    KnxDatapointType_DPT_DALI_Control_Gear_Diagnostic KnxDatapointType = 332
    KnxDatapointType_DPT_DALI_Diagnostics KnxDatapointType = 333
    KnxDatapointType_DPT_CombinedPosition KnxDatapointType = 334
    KnxDatapointType_DPT_StatusSAB KnxDatapointType = 335
    KnxDatapointType_DPT_Colour_xyY KnxDatapointType = 336
    KnxDatapointType_DPT_Converter_Status KnxDatapointType = 337
    KnxDatapointType_DPT_Converter_Test_Result KnxDatapointType = 338
    KnxDatapointType_DPT_Battery_Info KnxDatapointType = 339
    KnxDatapointType_DPT_Brightness_Colour_Temperature_Transition KnxDatapointType = 340
    KnxDatapointType_DPT_Brightness_Colour_Temperature_Control KnxDatapointType = 341
    KnxDatapointType_DPT_Colour_RGBW KnxDatapointType = 342
    KnxDatapointType_DPT_Relative_Control_RGBW KnxDatapointType = 343
    KnxDatapointType_DPT_Relative_Control_RGB KnxDatapointType = 344
    KnxDatapointType_DPT_GeographicalLocation KnxDatapointType = 345
    KnxDatapointType_DPT_TempRoomSetpSetF16_4 KnxDatapointType = 346
    KnxDatapointType_DPT_TempRoomSetpSetShiftF16_4 KnxDatapointType = 347
)


func (e KnxDatapointType) Number() uint16 {
    switch e  {
        case 0: { /* '0' */
            return 0
        }
        case 1: { /* '1' */
            return 0
        }
        case 10: { /* '10' */
            return 0
        }
        case 100: { /* '100' */
            return 2
        }
        case 101: { /* '101' */
            return 3
        }
        case 102: { /* '102' */
            return 4
        }
        case 103: { /* '103' */
            return 5
        }
        case 104: { /* '104' */
            return 6
        }
        case 105: { /* '105' */
            return 7
        }
        case 106: { /* '106' */
            return 8
        }
        case 107: { /* '107' */
            return 9
        }
        case 108: { /* '108' */
            return 10
        }
        case 109: { /* '109' */
            return 11
        }
        case 11: { /* '11' */
            return 0
        }
        case 110: { /* '110' */
            return 20
        }
        case 111: { /* '111' */
            return 21
        }
        case 112: { /* '112' */
            return 22
        }
        case 113: { /* '113' */
            return 23
        }
        case 114: { /* '114' */
            return 24
        }
        case 115: { /* '115' */
            return 25
        }
        case 116: { /* '116' */
            return 26
        }
        case 117: { /* '117' */
            return 27
        }
        case 118: { /* '118' */
            return 28
        }
        case 119: { /* '119' */
            return 29
        }
        case 12: { /* '12' */
            return 0
        }
        case 120: { /* '120' */
            return 30
        }
        case 121: { /* '121' */
            return 1
        }
        case 122: { /* '122' */
            return 1
        }
        case 123: { /* '123' */
            return 1
        }
        case 124: { /* '124' */
            return 100
        }
        case 125: { /* '125' */
            return 101
        }
        case 126: { /* '126' */
            return 102
        }
        case 127: { /* '127' */
            return 1200
        }
        case 128: { /* '128' */
            return 1201
        }
        case 129: { /* '129' */
            return 1
        }
        case 13: { /* '13' */
            return 0
        }
        case 130: { /* '130' */
            return 2
        }
        case 131: { /* '131' */
            return 10
        }
        case 132: { /* '132' */
            return 11
        }
        case 133: { /* '133' */
            return 12
        }
        case 134: { /* '134' */
            return 13
        }
        case 135: { /* '135' */
            return 14
        }
        case 136: { /* '136' */
            return 15
        }
        case 137: { /* '137' */
            return 16
        }
        case 138: { /* '138' */
            return 100
        }
        case 139: { /* '139' */
            return 1200
        }
        case 14: { /* '14' */
            return 0
        }
        case 140: { /* '140' */
            return 1201
        }
        case 141: { /* '141' */
            return 0
        }
        case 142: { /* '142' */
            return 1
        }
        case 143: { /* '143' */
            return 2
        }
        case 144: { /* '144' */
            return 3
        }
        case 145: { /* '145' */
            return 4
        }
        case 146: { /* '146' */
            return 5
        }
        case 147: { /* '147' */
            return 6
        }
        case 148: { /* '148' */
            return 7
        }
        case 149: { /* '149' */
            return 8
        }
        case 15: { /* '15' */
            return 0
        }
        case 150: { /* '150' */
            return 9
        }
        case 151: { /* '151' */
            return 10
        }
        case 152: { /* '152' */
            return 11
        }
        case 153: { /* '153' */
            return 12
        }
        case 154: { /* '154' */
            return 13
        }
        case 155: { /* '155' */
            return 14
        }
        case 156: { /* '156' */
            return 15
        }
        case 157: { /* '157' */
            return 16
        }
        case 158: { /* '158' */
            return 17
        }
        case 159: { /* '159' */
            return 18
        }
        case 16: { /* '16' */
            return 0
        }
        case 160: { /* '160' */
            return 19
        }
        case 161: { /* '161' */
            return 20
        }
        case 162: { /* '162' */
            return 21
        }
        case 163: { /* '163' */
            return 22
        }
        case 164: { /* '164' */
            return 23
        }
        case 165: { /* '165' */
            return 24
        }
        case 166: { /* '166' */
            return 25
        }
        case 167: { /* '167' */
            return 26
        }
        case 168: { /* '168' */
            return 27
        }
        case 169: { /* '169' */
            return 28
        }
        case 17: { /* '17' */
            return 0
        }
        case 170: { /* '170' */
            return 29
        }
        case 171: { /* '171' */
            return 30
        }
        case 172: { /* '172' */
            return 31
        }
        case 173: { /* '173' */
            return 32
        }
        case 174: { /* '174' */
            return 33
        }
        case 175: { /* '175' */
            return 34
        }
        case 176: { /* '176' */
            return 35
        }
        case 177: { /* '177' */
            return 36
        }
        case 178: { /* '178' */
            return 37
        }
        case 179: { /* '179' */
            return 38
        }
        case 18: { /* '18' */
            return 0
        }
        case 180: { /* '180' */
            return 39
        }
        case 181: { /* '181' */
            return 40
        }
        case 182: { /* '182' */
            return 41
        }
        case 183: { /* '183' */
            return 42
        }
        case 184: { /* '184' */
            return 43
        }
        case 185: { /* '185' */
            return 44
        }
        case 186: { /* '186' */
            return 45
        }
        case 187: { /* '187' */
            return 46
        }
        case 188: { /* '188' */
            return 47
        }
        case 189: { /* '189' */
            return 48
        }
        case 19: { /* '19' */
            return 0
        }
        case 190: { /* '190' */
            return 49
        }
        case 191: { /* '191' */
            return 50
        }
        case 192: { /* '192' */
            return 51
        }
        case 193: { /* '193' */
            return 52
        }
        case 194: { /* '194' */
            return 53
        }
        case 195: { /* '195' */
            return 54
        }
        case 196: { /* '196' */
            return 55
        }
        case 197: { /* '197' */
            return 56
        }
        case 198: { /* '198' */
            return 57
        }
        case 199: { /* '199' */
            return 58
        }
        case 2: { /* '2' */
            return 0
        }
        case 20: { /* '20' */
            return 0
        }
        case 200: { /* '200' */
            return 59
        }
        case 201: { /* '201' */
            return 60
        }
        case 202: { /* '202' */
            return 61
        }
        case 203: { /* '203' */
            return 62
        }
        case 204: { /* '204' */
            return 63
        }
        case 205: { /* '205' */
            return 64
        }
        case 206: { /* '206' */
            return 65
        }
        case 207: { /* '207' */
            return 66
        }
        case 208: { /* '208' */
            return 67
        }
        case 209: { /* '209' */
            return 68
        }
        case 21: { /* '21' */
            return 0
        }
        case 210: { /* '210' */
            return 69
        }
        case 211: { /* '211' */
            return 70
        }
        case 212: { /* '212' */
            return 71
        }
        case 213: { /* '213' */
            return 72
        }
        case 214: { /* '214' */
            return 73
        }
        case 215: { /* '215' */
            return 74
        }
        case 216: { /* '216' */
            return 75
        }
        case 217: { /* '217' */
            return 76
        }
        case 218: { /* '218' */
            return 77
        }
        case 219: { /* '219' */
            return 78
        }
        case 22: { /* '22' */
            return 0
        }
        case 220: { /* '220' */
            return 79
        }
        case 221: { /* '221' */
            return 1200
        }
        case 222: { /* '222' */
            return 1201
        }
        case 223: { /* '223' */
            return 0
        }
        case 224: { /* '224' */
            return 0
        }
        case 225: { /* '225' */
            return 1
        }
        case 226: { /* '226' */
            return 1
        }
        case 227: { /* '227' */
            return 1
        }
        case 228: { /* '228' */
            return 1
        }
        case 229: { /* '229' */
            return 1
        }
        case 23: { /* '23' */
            return 0
        }
        case 230: { /* '230' */
            return 2
        }
        case 231: { /* '231' */
            return 3
        }
        case 232: { /* '232' */
            return 4
        }
        case 233: { /* '233' */
            return 5
        }
        case 234: { /* '234' */
            return 6
        }
        case 235: { /* '235' */
            return 7
        }
        case 236: { /* '236' */
            return 8
        }
        case 237: { /* '237' */
            return 11
        }
        case 238: { /* '238' */
            return 12
        }
        case 239: { /* '239' */
            return 13
        }
        case 24: { /* '24' */
            return 0
        }
        case 240: { /* '240' */
            return 14
        }
        case 241: { /* '241' */
            return 17
        }
        case 242: { /* '242' */
            return 20
        }
        case 243: { /* '243' */
            return 21
        }
        case 244: { /* '244' */
            return 22
        }
        case 245: { /* '245' */
            return 100
        }
        case 246: { /* '246' */
            return 101
        }
        case 247: { /* '247' */
            return 102
        }
        case 248: { /* '248' */
            return 103
        }
        case 249: { /* '249' */
            return 104
        }
        case 25: { /* '25' */
            return 0
        }
        case 250: { /* '250' */
            return 105
        }
        case 251: { /* '251' */
            return 106
        }
        case 252: { /* '252' */
            return 107
        }
        case 253: { /* '253' */
            return 108
        }
        case 254: { /* '254' */
            return 109
        }
        case 255: { /* '255' */
            return 110
        }
        case 256: { /* '256' */
            return 111
        }
        case 257: { /* '257' */
            return 112
        }
        case 258: { /* '258' */
            return 113
        }
        case 259: { /* '259' */
            return 114
        }
        case 26: { /* '26' */
            return 0
        }
        case 260: { /* '260' */
            return 115
        }
        case 261: { /* '261' */
            return 116
        }
        case 262: { /* '262' */
            return 120
        }
        case 263: { /* '263' */
            return 121
        }
        case 264: { /* '264' */
            return 122
        }
        case 265: { /* '265' */
            return 600
        }
        case 266: { /* '266' */
            return 601
        }
        case 267: { /* '267' */
            return 602
        }
        case 268: { /* '268' */
            return 603
        }
        case 269: { /* '269' */
            return 604
        }
        case 27: { /* '27' */
            return 1
        }
        case 270: { /* '270' */
            return 605
        }
        case 271: { /* '271' */
            return 606
        }
        case 272: { /* '272' */
            return 607
        }
        case 273: { /* '273' */
            return 608
        }
        case 274: { /* '274' */
            return 609
        }
        case 275: { /* '275' */
            return 610
        }
        case 276: { /* '276' */
            return 611
        }
        case 277: { /* '277' */
            return 801
        }
        case 278: { /* '278' */
            return 802
        }
        case 279: { /* '279' */
            return 803
        }
        case 28: { /* '28' */
            return 2
        }
        case 280: { /* '280' */
            return 804
        }
        case 281: { /* '281' */
            return 1000
        }
        case 282: { /* '282' */
            return 1001
        }
        case 283: { /* '283' */
            return 1002
        }
        case 284: { /* '284' */
            return 1003
        }
        case 285: { /* '285' */
            return 1
        }
        case 286: { /* '286' */
            return 2
        }
        case 287: { /* '287' */
            return 100
        }
        case 288: { /* '288' */
            return 101
        }
        case 289: { /* '289' */
            return 102
        }
        case 29: { /* '29' */
            return 3
        }
        case 290: { /* '290' */
            return 103
        }
        case 291: { /* '291' */
            return 104
        }
        case 292: { /* '292' */
            return 105
        }
        case 293: { /* '293' */
            return 106
        }
        case 294: { /* '294' */
            return 107
        }
        case 295: { /* '295' */
            return 601
        }
        case 296: { /* '296' */
            return 1000
        }
        case 297: { /* '297' */
            return 1001
        }
        case 298: { /* '298' */
            return 1010
        }
        case 299: { /* '299' */
            return 100
        }
        case 3: { /* '3' */
            return 0
        }
        case 30: { /* '30' */
            return 4
        }
        case 300: { /* '300' */
            return 101
        }
        case 301: { /* '301' */
            return 102
        }
        case 302: { /* '302' */
            return 103
        }
        case 303: { /* '303' */
            return 1000
        }
        case 304: { /* '304' */
            return 1010
        }
        case 305: { /* '305' */
            return 1
        }
        case 306: { /* '306' */
            return 2
        }
        case 307: { /* '307' */
            return 3
        }
        case 308: { /* '308' */
            return 102
        }
        case 309: { /* '309' */
            return 1000
        }
        case 31: { /* '31' */
            return 5
        }
        case 310: { /* '310' */
            return 1
        }
        case 311: { /* '311' */
            return 1
        }
        case 312: { /* '312' */
            return 10
        }
        case 313: { /* '313' */
            return 11
        }
        case 314: { /* '314' */
            return 12
        }
        case 315: { /* '315' */
            return 1010
        }
        case 316: { /* '316' */
            return 100
        }
        case 317: { /* '317' */
            return 102
        }
        case 318: { /* '318' */
            return 104
        }
        case 319: { /* '319' */
            return 105
        }
        case 32: { /* '32' */
            return 6
        }
        case 320: { /* '320' */
            return 1
        }
        case 321: { /* '321' */
            return 1
        }
        case 322: { /* '322' */
            return 100
        }
        case 323: { /* '323' */
            return 101
        }
        case 324: { /* '324' */
            return 1
        }
        case 325: { /* '325' */
            return 2
        }
        case 326: { /* '326' */
            return 1
        }
        case 327: { /* '327' */
            return 1000
        }
        case 328: { /* '328' */
            return 600
        }
        case 329: { /* '329' */
            return 1
        }
        case 33: { /* '33' */
            return 7
        }
        case 330: { /* '330' */
            return 1
        }
        case 331: { /* '331' */
            return 1
        }
        case 332: { /* '332' */
            return 600
        }
        case 333: { /* '333' */
            return 600
        }
        case 334: { /* '334' */
            return 800
        }
        case 335: { /* '335' */
            return 800
        }
        case 336: { /* '336' */
            return 600
        }
        case 337: { /* '337' */
            return 600
        }
        case 338: { /* '338' */
            return 600
        }
        case 339: { /* '339' */
            return 600
        }
        case 34: { /* '34' */
            return 8
        }
        case 340: { /* '340' */
            return 600
        }
        case 341: { /* '341' */
            return 600
        }
        case 342: { /* '342' */
            return 600
        }
        case 343: { /* '343' */
            return 600
        }
        case 344: { /* '344' */
            return 600
        }
        case 345: { /* '345' */
            return 1
        }
        case 346: { /* '346' */
            return 100
        }
        case 347: { /* '347' */
            return 101
        }
        case 35: { /* '35' */
            return 9
        }
        case 36: { /* '36' */
            return 10
        }
        case 37: { /* '37' */
            return 11
        }
        case 38: { /* '38' */
            return 12
        }
        case 39: { /* '39' */
            return 13
        }
        case 4: { /* '4' */
            return 0
        }
        case 40: { /* '40' */
            return 14
        }
        case 41: { /* '41' */
            return 15
        }
        case 42: { /* '42' */
            return 16
        }
        case 43: { /* '43' */
            return 17
        }
        case 44: { /* '44' */
            return 18
        }
        case 45: { /* '45' */
            return 19
        }
        case 46: { /* '46' */
            return 21
        }
        case 47: { /* '47' */
            return 22
        }
        case 48: { /* '48' */
            return 23
        }
        case 49: { /* '49' */
            return 24
        }
        case 5: { /* '5' */
            return 0
        }
        case 50: { /* '50' */
            return 100
        }
        case 51: { /* '51' */
            return 1
        }
        case 52: { /* '52' */
            return 2
        }
        case 53: { /* '53' */
            return 3
        }
        case 54: { /* '54' */
            return 4
        }
        case 55: { /* '55' */
            return 5
        }
        case 56: { /* '56' */
            return 6
        }
        case 57: { /* '57' */
            return 7
        }
        case 58: { /* '58' */
            return 8
        }
        case 59: { /* '59' */
            return 9
        }
        case 6: { /* '6' */
            return 0
        }
        case 60: { /* '60' */
            return 10
        }
        case 61: { /* '61' */
            return 11
        }
        case 62: { /* '62' */
            return 12
        }
        case 63: { /* '63' */
            return 7
        }
        case 64: { /* '64' */
            return 8
        }
        case 65: { /* '65' */
            return 1
        }
        case 66: { /* '66' */
            return 2
        }
        case 67: { /* '67' */
            return 1
        }
        case 68: { /* '68' */
            return 3
        }
        case 69: { /* '69' */
            return 4
        }
        case 7: { /* '7' */
            return 0
        }
        case 70: { /* '70' */
            return 5
        }
        case 71: { /* '71' */
            return 6
        }
        case 72: { /* '72' */
            return 10
        }
        case 73: { /* '73' */
            return 100
        }
        case 74: { /* '74' */
            return 1
        }
        case 75: { /* '75' */
            return 10
        }
        case 76: { /* '76' */
            return 20
        }
        case 77: { /* '77' */
            return 1
        }
        case 78: { /* '78' */
            return 2
        }
        case 79: { /* '79' */
            return 3
        }
        case 8: { /* '8' */
            return 0
        }
        case 80: { /* '80' */
            return 4
        }
        case 81: { /* '81' */
            return 5
        }
        case 82: { /* '82' */
            return 6
        }
        case 83: { /* '83' */
            return 7
        }
        case 84: { /* '84' */
            return 10
        }
        case 85: { /* '85' */
            return 11
        }
        case 86: { /* '86' */
            return 12
        }
        case 87: { /* '87' */
            return 13
        }
        case 88: { /* '88' */
            return 600
        }
        case 89: { /* '89' */
            return 1
        }
        case 9: { /* '9' */
            return 0
        }
        case 90: { /* '90' */
            return 2
        }
        case 91: { /* '91' */
            return 3
        }
        case 92: { /* '92' */
            return 4
        }
        case 93: { /* '93' */
            return 5
        }
        case 94: { /* '94' */
            return 6
        }
        case 95: { /* '95' */
            return 7
        }
        case 96: { /* '96' */
            return 10
        }
        case 97: { /* '97' */
            return 11
        }
        case 98: { /* '98' */
            return 12
        }
        case 99: { /* '99' */
            return 1
        }
        default: {
            return 0
        }
    }
}

func (e KnxDatapointType) Name() string {
    switch e  {
        case 0: { /* '0' */
            return "Unknown Datapoint Subtype"
        }
        case 1: { /* '1' */
            return "BOOL"
        }
        case 10: { /* '10' */
            return "UDINT"
        }
        case 100: { /* '100' */
            return "temperature difference (K)"
        }
        case 101: { /* '101' */
            return "kelvin/hour (K/h)"
        }
        case 102: { /* '102' */
            return "lux (Lux)"
        }
        case 103: { /* '103' */
            return "speed (m/s)"
        }
        case 104: { /* '104' */
            return "pressure (Pa)"
        }
        case 105: { /* '105' */
            return "humidity (%)"
        }
        case 106: { /* '106' */
            return "parts/million (ppm)"
        }
        case 107: { /* '107' */
            return "air flow (m³/h)"
        }
        case 108: { /* '108' */
            return "time (s)"
        }
        case 109: { /* '109' */
            return "time (ms)"
        }
        case 11: { /* '11' */
            return "DINT"
        }
        case 110: { /* '110' */
            return "voltage (mV)"
        }
        case 111: { /* '111' */
            return "current (mA)"
        }
        case 112: { /* '112' */
            return "power denisity (W/m²)"
        }
        case 113: { /* '113' */
            return "kelvin/percent (K/%)"
        }
        case 114: { /* '114' */
            return "power (kW)"
        }
        case 115: { /* '115' */
            return "volume flow (l/h)"
        }
        case 116: { /* '116' */
            return "rain amount (l/m²)"
        }
        case 117: { /* '117' */
            return "temperature (°F)"
        }
        case 118: { /* '118' */
            return "wind speed (km/h)"
        }
        case 119: { /* '119' */
            return "absolute humidity (g/m³)"
        }
        case 12: { /* '12' */
            return "ULINT"
        }
        case 120: { /* '120' */
            return "concentration (µg/m³)"
        }
        case 121: { /* '121' */
            return "time of day"
        }
        case 122: { /* '122' */
            return "date"
        }
        case 123: { /* '123' */
            return "counter pulses (unsigned)"
        }
        case 124: { /* '124' */
            return "counter timesec (s)"
        }
        case 125: { /* '125' */
            return "counter timemin (min)"
        }
        case 126: { /* '126' */
            return "counter timehrs (h)"
        }
        case 127: { /* '127' */
            return "volume liquid (l)"
        }
        case 128: { /* '128' */
            return "volume (m³)"
        }
        case 129: { /* '129' */
            return "counter pulses (signed)"
        }
        case 13: { /* '13' */
            return "LINT"
        }
        case 130: { /* '130' */
            return "flow rate (m³/h)"
        }
        case 131: { /* '131' */
            return "active energy (Wh)"
        }
        case 132: { /* '132' */
            return "apparant energy (VAh)"
        }
        case 133: { /* '133' */
            return "reactive energy (VARh)"
        }
        case 134: { /* '134' */
            return "active energy (kWh)"
        }
        case 135: { /* '135' */
            return "apparant energy (kVAh)"
        }
        case 136: { /* '136' */
            return "reactive energy (kVARh)"
        }
        case 137: { /* '137' */
            return "active energy (MWh)"
        }
        case 138: { /* '138' */
            return "time lag (s)"
        }
        case 139: { /* '139' */
            return "delta volume liquid (l)"
        }
        case 14: { /* '14' */
            return "REAL"
        }
        case 140: { /* '140' */
            return "delta volume (m³)"
        }
        case 141: { /* '141' */
            return "acceleration (m/s²)"
        }
        case 142: { /* '142' */
            return "angular acceleration (rad/s²)"
        }
        case 143: { /* '143' */
            return "activation energy (J/mol)"
        }
        case 144: { /* '144' */
            return "radioactive activity (1/s)"
        }
        case 145: { /* '145' */
            return "amount of substance (mol)"
        }
        case 146: { /* '146' */
            return "amplitude"
        }
        case 147: { /* '147' */
            return "angle (radiant)"
        }
        case 148: { /* '148' */
            return "angle (degree)"
        }
        case 149: { /* '149' */
            return "angular momentum (Js)"
        }
        case 15: { /* '15' */
            return "LREAL"
        }
        case 150: { /* '150' */
            return "angular velocity (rad/s)"
        }
        case 151: { /* '151' */
            return "area (m*m)"
        }
        case 152: { /* '152' */
            return "capacitance (F)"
        }
        case 153: { /* '153' */
            return "flux density (C/m²)"
        }
        case 154: { /* '154' */
            return "charge density (C/m³)"
        }
        case 155: { /* '155' */
            return "compressibility (m²/N)"
        }
        case 156: { /* '156' */
            return "conductance (S)"
        }
        case 157: { /* '157' */
            return "conductivity  (S/m)"
        }
        case 158: { /* '158' */
            return "density (kg/m³)"
        }
        case 159: { /* '159' */
            return "electric charge (C)"
        }
        case 16: { /* '16' */
            return "CHAR"
        }
        case 160: { /* '160' */
            return "electric current (A)"
        }
        case 161: { /* '161' */
            return "electric current density (A/m²)"
        }
        case 162: { /* '162' */
            return "electric dipole moment (Cm)"
        }
        case 163: { /* '163' */
            return "electric displacement (C/m²)"
        }
        case 164: { /* '164' */
            return "electric field strength (V/m)"
        }
        case 165: { /* '165' */
            return "electric flux (C)"
        }
        case 166: { /* '166' */
            return "electric flux density (C/m²)"
        }
        case 167: { /* '167' */
            return "electric polarization (C/m²)"
        }
        case 168: { /* '168' */
            return "electric potential (V)"
        }
        case 169: { /* '169' */
            return "electric potential difference (V)"
        }
        case 17: { /* '17' */
            return "WCHAR"
        }
        case 170: { /* '170' */
            return "electromagnetic moment (Am²)"
        }
        case 171: { /* '171' */
            return "electromotive force (V)"
        }
        case 172: { /* '172' */
            return "energy (J)"
        }
        case 173: { /* '173' */
            return "force (N)"
        }
        case 174: { /* '174' */
            return "frequency (Hz)"
        }
        case 175: { /* '175' */
            return "angular frequency (rad/s)"
        }
        case 176: { /* '176' */
            return "heat capacity (J/K)"
        }
        case 177: { /* '177' */
            return "heat flow rate (W)"
        }
        case 178: { /* '178' */
            return "heat quantity"
        }
        case 179: { /* '179' */
            return "impedance (Ω)"
        }
        case 18: { /* '18' */
            return "STRING"
        }
        case 180: { /* '180' */
            return "length (m)"
        }
        case 181: { /* '181' */
            return "light quantity (J)"
        }
        case 182: { /* '182' */
            return "luminance (cd/m²)"
        }
        case 183: { /* '183' */
            return "luminous flux (lm)"
        }
        case 184: { /* '184' */
            return "luminous intensity (cd)"
        }
        case 185: { /* '185' */
            return "magnetic field strength (A/m)"
        }
        case 186: { /* '186' */
            return "magnetic flux (Wb)"
        }
        case 187: { /* '187' */
            return "magnetic flux density (T)"
        }
        case 188: { /* '188' */
            return "magnetic moment (Am²)"
        }
        case 189: { /* '189' */
            return "magnetic polarization (T)"
        }
        case 19: { /* '19' */
            return "WSTRING"
        }
        case 190: { /* '190' */
            return "magnetization (A/m)"
        }
        case 191: { /* '191' */
            return "magnetomotive force (A)"
        }
        case 192: { /* '192' */
            return "mass (kg)"
        }
        case 193: { /* '193' */
            return "mass flux (kg/s)"
        }
        case 194: { /* '194' */
            return "momentum (N/s)"
        }
        case 195: { /* '195' */
            return "phase angle (rad)"
        }
        case 196: { /* '196' */
            return "phase angle (°)"
        }
        case 197: { /* '197' */
            return "power (W)"
        }
        case 198: { /* '198' */
            return "power factor (cos Φ)"
        }
        case 199: { /* '199' */
            return "pressure (Pa)"
        }
        case 2: { /* '2' */
            return "BYTE"
        }
        case 20: { /* '20' */
            return "TIME"
        }
        case 200: { /* '200' */
            return "reactance (Ω)"
        }
        case 201: { /* '201' */
            return "resistance (Ω)"
        }
        case 202: { /* '202' */
            return "resistivity (Ωm)"
        }
        case 203: { /* '203' */
            return "self inductance (H)"
        }
        case 204: { /* '204' */
            return "solid angle (sr)"
        }
        case 205: { /* '205' */
            return "sound intensity (W/m²)"
        }
        case 206: { /* '206' */
            return "speed (m/s)"
        }
        case 207: { /* '207' */
            return "stress (Pa)"
        }
        case 208: { /* '208' */
            return "surface tension (N/m)"
        }
        case 209: { /* '209' */
            return "temperature (°C)"
        }
        case 21: { /* '21' */
            return "LTIME"
        }
        case 210: { /* '210' */
            return "temperature absolute (K)"
        }
        case 211: { /* '211' */
            return "temperature difference (K)"
        }
        case 212: { /* '212' */
            return "thermal capacity (J/K)"
        }
        case 213: { /* '213' */
            return "thermal conductivity (W/mK)"
        }
        case 214: { /* '214' */
            return "thermoelectric power (V/K)"
        }
        case 215: { /* '215' */
            return "time (s)"
        }
        case 216: { /* '216' */
            return "torque (Nm)"
        }
        case 217: { /* '217' */
            return "volume (m³)"
        }
        case 218: { /* '218' */
            return "volume flux (m³/s)"
        }
        case 219: { /* '219' */
            return "weight (N)"
        }
        case 22: { /* '22' */
            return "DATE"
        }
        case 220: { /* '220' */
            return "work (J)"
        }
        case 221: { /* '221' */
            return "volume flux for meters (m³/h)"
        }
        case 222: { /* '222' */
            return "volume flux for meters (1/ls)"
        }
        case 223: { /* '223' */
            return "entrance access"
        }
        case 224: { /* '224' */
            return "Character String (ASCII)"
        }
        case 225: { /* '225' */
            return "Character String (ISO 8859-1)"
        }
        case 226: { /* '226' */
            return "scene number"
        }
        case 227: { /* '227' */
            return "scene control"
        }
        case 228: { /* '228' */
            return "date time"
        }
        case 229: { /* '229' */
            return "SCLO mode"
        }
        case 23: { /* '23' */
            return "TIME_OF_DAY"
        }
        case 230: { /* '230' */
            return "building mode"
        }
        case 231: { /* '231' */
            return "occupied"
        }
        case 232: { /* '232' */
            return "priority"
        }
        case 233: { /* '233' */
            return "light application mode"
        }
        case 234: { /* '234' */
            return "light application area"
        }
        case 235: { /* '235' */
            return "alarm class"
        }
        case 236: { /* '236' */
            return "PSU mode"
        }
        case 237: { /* '237' */
            return "system error class"
        }
        case 238: { /* '238' */
            return "HVAC error class"
        }
        case 239: { /* '239' */
            return "time delay"
        }
        case 24: { /* '24' */
            return "TOD"
        }
        case 240: { /* '240' */
            return "wind force scale (0..12)"
        }
        case 241: { /* '241' */
            return "sensor mode"
        }
        case 242: { /* '242' */
            return "actuator connect type"
        }
        case 243: { /* '243' */
            return "cloud cover"
        }
        case 244: { /* '244' */
            return "power return mode"
        }
        case 245: { /* '245' */
            return "fuel type"
        }
        case 246: { /* '246' */
            return "burner type"
        }
        case 247: { /* '247' */
            return "HVAC mode"
        }
        case 248: { /* '248' */
            return "DHW mode"
        }
        case 249: { /* '249' */
            return "load priority"
        }
        case 25: { /* '25' */
            return "DATE_AND_TIME"
        }
        case 250: { /* '250' */
            return "HVAC control mode"
        }
        case 251: { /* '251' */
            return "HVAC emergency mode"
        }
        case 252: { /* '252' */
            return "changeover mode"
        }
        case 253: { /* '253' */
            return "valve mode"
        }
        case 254: { /* '254' */
            return "damper mode"
        }
        case 255: { /* '255' */
            return "heater mode"
        }
        case 256: { /* '256' */
            return "fan mode"
        }
        case 257: { /* '257' */
            return "master/slave mode"
        }
        case 258: { /* '258' */
            return "status room setpoint"
        }
        case 259: { /* '259' */
            return "metering device type"
        }
        case 26: { /* '26' */
            return "DT"
        }
        case 260: { /* '260' */
            return "hum dehum mode"
        }
        case 261: { /* '261' */
            return "enable H/C stage"
        }
        case 262: { /* '262' */
            return "ADA type"
        }
        case 263: { /* '263' */
            return "backup mode"
        }
        case 264: { /* '264' */
            return "start syncronization type"
        }
        case 265: { /* '265' */
            return "behavior lock/unlock"
        }
        case 266: { /* '266' */
            return "behavior bus power up/down"
        }
        case 267: { /* '267' */
            return "dali fade time"
        }
        case 268: { /* '268' */
            return "blink mode"
        }
        case 269: { /* '269' */
            return "light control mode"
        }
        case 27: { /* '27' */
            return "switch"
        }
        case 270: { /* '270' */
            return "PB switch mode"
        }
        case 271: { /* '271' */
            return "PB action mode"
        }
        case 272: { /* '272' */
            return "PB dimm mode"
        }
        case 273: { /* '273' */
            return "switch on mode"
        }
        case 274: { /* '274' */
            return "load type"
        }
        case 275: { /* '275' */
            return "load type detection"
        }
        case 276: { /* '276' */
            return "converter test control"
        }
        case 277: { /* '277' */
            return "SAB except behavior"
        }
        case 278: { /* '278' */
            return "SAB behavior on lock/unlock"
        }
        case 279: { /* '279' */
            return "SSSB mode"
        }
        case 28: { /* '28' */
            return "boolean"
        }
        case 280: { /* '280' */
            return "blinds control mode"
        }
        case 281: { /* '281' */
            return "communication mode"
        }
        case 282: { /* '282' */
            return "additional information type"
        }
        case 283: { /* '283' */
            return "RF mode selection"
        }
        case 284: { /* '284' */
            return "RF filter mode selection"
        }
        case 285: { /* '285' */
            return "general status"
        }
        case 286: { /* '286' */
            return "device control"
        }
        case 287: { /* '287' */
            return "forcing signal"
        }
        case 288: { /* '288' */
            return "forcing signal cool"
        }
        case 289: { /* '289' */
            return "room heating controller status"
        }
        case 29: { /* '29' */
            return "enable"
        }
        case 290: { /* '290' */
            return "solar DHW controller status"
        }
        case 291: { /* '291' */
            return "fuel type set"
        }
        case 292: { /* '292' */
            return "room cooling controller status"
        }
        case 293: { /* '293' */
            return "ventilation controller status"
        }
        case 294: { /* '294' */
            return "combined status RTSM"
        }
        case 295: { /* '295' */
            return "lighting actuator error information"
        }
        case 296: { /* '296' */
            return "RF communication mode info"
        }
        case 297: { /* '297' */
            return "cEMI server supported RF filtering modes"
        }
        case 298: { /* '298' */
            return "channel activation for 8 channels"
        }
        case 299: { /* '299' */
            return "DHW controller status"
        }
        case 3: { /* '3' */
            return "WORD"
        }
        case 30: { /* '30' */
            return "ramp"
        }
        case 300: { /* '300' */
            return "RHCC status"
        }
        case 301: { /* '301' */
            return "combined status HVA"
        }
        case 302: { /* '302' */
            return "combined status RTC"
        }
        case 303: { /* '303' */
            return "media"
        }
        case 304: { /* '304' */
            return "channel activation for 16 channels"
        }
        case 305: { /* '305' */
            return "on/off action"
        }
        case 306: { /* '306' */
            return "alarm reaction"
        }
        case 307: { /* '307' */
            return "up/down action"
        }
        case 308: { /* '308' */
            return "HVAC push button action"
        }
        case 309: { /* '309' */
            return "busy/nak repetitions"
        }
        case 31: { /* '31' */
            return "alarm"
        }
        case 310: { /* '310' */
            return "scene information"
        }
        case 311: { /* '311' */
            return "bit-combined info on/off"
        }
        case 312: { /* '312' */
            return "active energy (Wh)"
        }
        case 313: { /* '313' */
            return "apparant energy (VAh)"
        }
        case 314: { /* '314' */
            return "reactive energy (VARh)"
        }
        case 315: { /* '315' */
            return "activation state 0..23"
        }
        case 316: { /* '316' */
            return "time delay & HVAC mode"
        }
        case 317: { /* '317' */
            return "time delay & DHW mode"
        }
        case 318: { /* '318' */
            return "time delay & occupancy mode"
        }
        case 319: { /* '319' */
            return "time delay & building mode"
        }
        case 32: { /* '32' */
            return "binary value"
        }
        case 320: { /* '320' */
            return "DPT version"
        }
        case 321: { /* '321' */
            return "alarm info"
        }
        case 322: { /* '322' */
            return "room temperature setpoint"
        }
        case 323: { /* '323' */
            return "room temperature setpoint shift"
        }
        case 324: { /* '324' */
            return "scaling speed"
        }
        case 325: { /* '325' */
            return "scaling step time"
        }
        case 326: { /* '326' */
            return "metering value (value,encoding,cmd)"
        }
        case 327: { /* '327' */
            return "MBus address"
        }
        case 328: { /* '328' */
            return "RGB value 3x(0..255)"
        }
        case 329: { /* '329' */
            return "language code (ASCII)"
        }
        case 33: { /* '33' */
            return "step"
        }
        case 330: { /* '330' */
            return "electrical energy with tariff"
        }
        case 331: { /* '331' */
            return "priority control"
        }
        case 332: { /* '332' */
            return "diagnostic value"
        }
        case 333: { /* '333' */
            return "diagnostic value"
        }
        case 334: { /* '334' */
            return "combined position"
        }
        case 335: { /* '335' */
            return "status sunblind & shutter actuator"
        }
        case 336: { /* '336' */
            return "colour xyY"
        }
        case 337: { /* '337' */
            return "DALI converter status"
        }
        case 338: { /* '338' */
            return "DALI converter test result"
        }
        case 339: { /* '339' */
            return "Battery Information"
        }
        case 34: { /* '34' */
            return "up/down"
        }
        case 340: { /* '340' */
            return "brightness colour temperature transition"
        }
        case 341: { /* '341' */
            return "brightness colour temperature control"
        }
        case 342: { /* '342' */
            return "RGBW value 4x(0..100%)"
        }
        case 343: { /* '343' */
            return "RGBW relative control"
        }
        case 344: { /* '344' */
            return "RGB relative control"
        }
        case 345: { /* '345' */
            return "geographical location (longitude and latitude) expressed in degrees"
        }
        case 346: { /* '346' */
            return "Temperature setpoint setting for 4 HVAC Modes"
        }
        case 347: { /* '347' */
            return "Temperature setpoint shift setting for 4 HVAC Modes"
        }
        case 35: { /* '35' */
            return "open/close"
        }
        case 36: { /* '36' */
            return "start/stop"
        }
        case 37: { /* '37' */
            return "state"
        }
        case 38: { /* '38' */
            return "invert"
        }
        case 39: { /* '39' */
            return "dim send style"
        }
        case 4: { /* '4' */
            return "DWORD"
        }
        case 40: { /* '40' */
            return "input source"
        }
        case 41: { /* '41' */
            return "reset"
        }
        case 42: { /* '42' */
            return "acknowledge"
        }
        case 43: { /* '43' */
            return "trigger"
        }
        case 44: { /* '44' */
            return "occupancy"
        }
        case 45: { /* '45' */
            return "window/door"
        }
        case 46: { /* '46' */
            return "logical function"
        }
        case 47: { /* '47' */
            return "scene"
        }
        case 48: { /* '48' */
            return "shutter/blinds mode"
        }
        case 49: { /* '49' */
            return "day/night"
        }
        case 5: { /* '5' */
            return "LWORD"
        }
        case 50: { /* '50' */
            return "cooling/heating"
        }
        case 51: { /* '51' */
            return "switch control"
        }
        case 52: { /* '52' */
            return "boolean control"
        }
        case 53: { /* '53' */
            return "enable control"
        }
        case 54: { /* '54' */
            return "ramp control"
        }
        case 55: { /* '55' */
            return "alarm control"
        }
        case 56: { /* '56' */
            return "binary value control"
        }
        case 57: { /* '57' */
            return "step control"
        }
        case 58: { /* '58' */
            return "direction control 1"
        }
        case 59: { /* '59' */
            return "direction control 2"
        }
        case 6: { /* '6' */
            return "USINT"
        }
        case 60: { /* '60' */
            return "start control"
        }
        case 61: { /* '61' */
            return "state control"
        }
        case 62: { /* '62' */
            return "invert control"
        }
        case 63: { /* '63' */
            return "dimming control"
        }
        case 64: { /* '64' */
            return "blind control"
        }
        case 65: { /* '65' */
            return "character (ASCII)"
        }
        case 66: { /* '66' */
            return "character (ISO 8859-1)"
        }
        case 67: { /* '67' */
            return "percentage (0..100%)"
        }
        case 68: { /* '68' */
            return "angle (degrees)"
        }
        case 69: { /* '69' */
            return "percentage (0..255%)"
        }
        case 7: { /* '7' */
            return "SINT"
        }
        case 70: { /* '70' */
            return "ratio (0..255)"
        }
        case 71: { /* '71' */
            return "tariff (0..255)"
        }
        case 72: { /* '72' */
            return "counter pulses (0..255)"
        }
        case 73: { /* '73' */
            return "fan stage (0..255)"
        }
        case 74: { /* '74' */
            return "percentage (-128..127%)"
        }
        case 75: { /* '75' */
            return "counter pulses (-128..127)"
        }
        case 76: { /* '76' */
            return "status with mode"
        }
        case 77: { /* '77' */
            return "pulses"
        }
        case 78: { /* '78' */
            return "time (ms)"
        }
        case 79: { /* '79' */
            return "time (10 ms)"
        }
        case 8: { /* '8' */
            return "UINT"
        }
        case 80: { /* '80' */
            return "time (100 ms)"
        }
        case 81: { /* '81' */
            return "time (s)"
        }
        case 82: { /* '82' */
            return "time (min)"
        }
        case 83: { /* '83' */
            return "time (h)"
        }
        case 84: { /* '84' */
            return "property data type"
        }
        case 85: { /* '85' */
            return "length (mm)"
        }
        case 86: { /* '86' */
            return "current (mA)"
        }
        case 87: { /* '87' */
            return "brightness (lux)"
        }
        case 88: { /* '88' */
            return "absolute colour temperature (K)"
        }
        case 89: { /* '89' */
            return "pulses difference"
        }
        case 9: { /* '9' */
            return "INT"
        }
        case 90: { /* '90' */
            return "time lag (ms)"
        }
        case 91: { /* '91' */
            return "time lag(10 ms)"
        }
        case 92: { /* '92' */
            return "time lag (100 ms)"
        }
        case 93: { /* '93' */
            return "time lag (s)"
        }
        case 94: { /* '94' */
            return "time lag (min)"
        }
        case 95: { /* '95' */
            return "time lag (h)"
        }
        case 96: { /* '96' */
            return "percentage difference (%)"
        }
        case 97: { /* '97' */
            return "rotation angle (°)"
        }
        case 98: { /* '98' */
            return "length (m)"
        }
        case 99: { /* '99' */
            return "temperature (°C)"
        }
        default: {
            return ""
        }
    }
}

func (e KnxDatapointType) DatapointMainType() KnxDatapointMainType {
    switch e  {
        case 0: { /* '0' */
            return KnxDatapointMainType_DPT_UNKNOWN
        }
        case 1: { /* '1' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 10: { /* '10' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 100: { /* '100' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 101: { /* '101' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 102: { /* '102' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 103: { /* '103' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 104: { /* '104' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 105: { /* '105' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 106: { /* '106' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 107: { /* '107' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 108: { /* '108' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 109: { /* '109' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 11: { /* '11' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 110: { /* '110' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 111: { /* '111' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 112: { /* '112' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 113: { /* '113' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 114: { /* '114' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 115: { /* '115' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 116: { /* '116' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 117: { /* '117' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 118: { /* '118' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 119: { /* '119' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 12: { /* '12' */
            return KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE
        }
        case 120: { /* '120' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        case 121: { /* '121' */
            return KnxDatapointMainType_DPT_TIME
        }
        case 122: { /* '122' */
            return KnxDatapointMainType_DPT_DATE
        }
        case 123: { /* '123' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 124: { /* '124' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 125: { /* '125' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 126: { /* '126' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 127: { /* '127' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 128: { /* '128' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 129: { /* '129' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 13: { /* '13' */
            return KnxDatapointMainType_DPT_8_BYTE_SIGNED_VALUE
        }
        case 130: { /* '130' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 131: { /* '131' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 132: { /* '132' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 133: { /* '133' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 134: { /* '134' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 135: { /* '135' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 136: { /* '136' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 137: { /* '137' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 138: { /* '138' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 139: { /* '139' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 14: { /* '14' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 140: { /* '140' */
            return KnxDatapointMainType_DPT_4_BYTE_SIGNED_VALUE
        }
        case 141: { /* '141' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 142: { /* '142' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 143: { /* '143' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 144: { /* '144' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 145: { /* '145' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 146: { /* '146' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 147: { /* '147' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 148: { /* '148' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 149: { /* '149' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 15: { /* '15' */
            return KnxDatapointMainType_DPT_8_BYTE_FLOAT_VALUE
        }
        case 150: { /* '150' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 151: { /* '151' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 152: { /* '152' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 153: { /* '153' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 154: { /* '154' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 155: { /* '155' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 156: { /* '156' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 157: { /* '157' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 158: { /* '158' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 159: { /* '159' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 16: { /* '16' */
            return KnxDatapointMainType_DPT_CHARACTER
        }
        case 160: { /* '160' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 161: { /* '161' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 162: { /* '162' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 163: { /* '163' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 164: { /* '164' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 165: { /* '165' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 166: { /* '166' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 167: { /* '167' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 168: { /* '168' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 169: { /* '169' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 17: { /* '17' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 170: { /* '170' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 171: { /* '171' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 172: { /* '172' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 173: { /* '173' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 174: { /* '174' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 175: { /* '175' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 176: { /* '176' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 177: { /* '177' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 178: { /* '178' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 179: { /* '179' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 18: { /* '18' */
            return KnxDatapointMainType_DPT_UNKNOWN
        }
        case 180: { /* '180' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 181: { /* '181' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 182: { /* '182' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 183: { /* '183' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 184: { /* '184' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 185: { /* '185' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 186: { /* '186' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 187: { /* '187' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 188: { /* '188' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 189: { /* '189' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 19: { /* '19' */
            return KnxDatapointMainType_DPT_UNKNOWN
        }
        case 190: { /* '190' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 191: { /* '191' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 192: { /* '192' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 193: { /* '193' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 194: { /* '194' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 195: { /* '195' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 196: { /* '196' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 197: { /* '197' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 198: { /* '198' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 199: { /* '199' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 2: { /* '2' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 20: { /* '20' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 200: { /* '200' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 201: { /* '201' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 202: { /* '202' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 203: { /* '203' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 204: { /* '204' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 205: { /* '205' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 206: { /* '206' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 207: { /* '207' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 208: { /* '208' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 209: { /* '209' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 21: { /* '21' */
            return KnxDatapointMainType_DPT_8_BYTE_UNSIGNED_VALUE
        }
        case 210: { /* '210' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 211: { /* '211' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 212: { /* '212' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 213: { /* '213' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 214: { /* '214' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 215: { /* '215' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 216: { /* '216' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 217: { /* '217' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 218: { /* '218' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 219: { /* '219' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 22: { /* '22' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 220: { /* '220' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 221: { /* '221' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 222: { /* '222' */
            return KnxDatapointMainType_DPT_4_BYTE_FLOAT_VALUE
        }
        case 223: { /* '223' */
            return KnxDatapointMainType_DPT_ENTRANCE_ACCESS
        }
        case 224: { /* '224' */
            return KnxDatapointMainType_DPT_CHARACTER_STRING
        }
        case 225: { /* '225' */
            return KnxDatapointMainType_DPT_CHARACTER_STRING
        }
        case 226: { /* '226' */
            return KnxDatapointMainType_DPT_SCENE_NUMBER
        }
        case 227: { /* '227' */
            return KnxDatapointMainType_DPT_SCENE_CONTROL
        }
        case 228: { /* '228' */
            return KnxDatapointMainType_DPT_DATE_TIME
        }
        case 229: { /* '229' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 23: { /* '23' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 230: { /* '230' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 231: { /* '231' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 232: { /* '232' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 233: { /* '233' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 234: { /* '234' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 235: { /* '235' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 236: { /* '236' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 237: { /* '237' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 238: { /* '238' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 239: { /* '239' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 24: { /* '24' */
            return KnxDatapointMainType_DPT_4_BYTE_UNSIGNED_VALUE
        }
        case 240: { /* '240' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 241: { /* '241' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 242: { /* '242' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 243: { /* '243' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 244: { /* '244' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 245: { /* '245' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 246: { /* '246' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 247: { /* '247' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 248: { /* '248' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 249: { /* '249' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 25: { /* '25' */
            return KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE
        }
        case 250: { /* '250' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 251: { /* '251' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 252: { /* '252' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 253: { /* '253' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 254: { /* '254' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 255: { /* '255' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 256: { /* '256' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 257: { /* '257' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 258: { /* '258' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 259: { /* '259' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 26: { /* '26' */
            return KnxDatapointMainType_DPT_12_BYTE_SIGNED_VALUE
        }
        case 260: { /* '260' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 261: { /* '261' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 262: { /* '262' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 263: { /* '263' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 264: { /* '264' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 265: { /* '265' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 266: { /* '266' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 267: { /* '267' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 268: { /* '268' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 269: { /* '269' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 27: { /* '27' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 270: { /* '270' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 271: { /* '271' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 272: { /* '272' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 273: { /* '273' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 274: { /* '274' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 275: { /* '275' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 276: { /* '276' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 277: { /* '277' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 278: { /* '278' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 279: { /* '279' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 28: { /* '28' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 280: { /* '280' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 281: { /* '281' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 282: { /* '282' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 283: { /* '283' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 284: { /* '284' */
            return KnxDatapointMainType_DPT_1_BYTE
        }
        case 285: { /* '285' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 286: { /* '286' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 287: { /* '287' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 288: { /* '288' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 289: { /* '289' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 29: { /* '29' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 290: { /* '290' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 291: { /* '291' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 292: { /* '292' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 293: { /* '293' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 294: { /* '294' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 295: { /* '295' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 296: { /* '296' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 297: { /* '297' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 298: { /* '298' */
            return KnxDatapointMainType_DPT_8_BIT_SET
        }
        case 299: { /* '299' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 3: { /* '3' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 30: { /* '30' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 300: { /* '300' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 301: { /* '301' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 302: { /* '302' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 303: { /* '303' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 304: { /* '304' */
            return KnxDatapointMainType_DPT_16_BIT_SET
        }
        case 305: { /* '305' */
            return KnxDatapointMainType_DPT_2_BIT_SET
        }
        case 306: { /* '306' */
            return KnxDatapointMainType_DPT_2_BIT_SET
        }
        case 307: { /* '307' */
            return KnxDatapointMainType_DPT_2_BIT_SET
        }
        case 308: { /* '308' */
            return KnxDatapointMainType_DPT_2_BIT_SET
        }
        case 309: { /* '309' */
            return KnxDatapointMainType_DPT_2_NIBBLE_SET
        }
        case 31: { /* '31' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 310: { /* '310' */
            return KnxDatapointMainType_DPT_8_BIT_SET_2
        }
        case 311: { /* '311' */
            return KnxDatapointMainType_DPT_32_BIT_SET
        }
        case 312: { /* '312' */
            return KnxDatapointMainType_DPT_ELECTRICAL_ENERGY
        }
        case 313: { /* '313' */
            return KnxDatapointMainType_DPT_ELECTRICAL_ENERGY
        }
        case 314: { /* '314' */
            return KnxDatapointMainType_DPT_ELECTRICAL_ENERGY
        }
        case 315: { /* '315' */
            return KnxDatapointMainType_DPT_24_TIMES_CHANNEL_ACTIVATION
        }
        case 316: { /* '316' */
            return KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
        }
        case 317: { /* '317' */
            return KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
        }
        case 318: { /* '318' */
            return KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
        }
        case 319: { /* '319' */
            return KnxDatapointMainType_DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM
        }
        case 32: { /* '32' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 320: { /* '320' */
            return KnxDatapointMainType_DPT_DATAPOINT_TYPE_VERSION
        }
        case 321: { /* '321' */
            return KnxDatapointMainType_DPT_ALARM_INFO
        }
        case 322: { /* '322' */
            return KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE
        }
        case 323: { /* '323' */
            return KnxDatapointMainType_DPT_3X_2_BYTE_FLOAT_VALUE
        }
        case 324: { /* '324' */
            return KnxDatapointMainType_DPT_SCALING_SPEED
        }
        case 325: { /* '325' */
            return KnxDatapointMainType_DPT_SCALING_SPEED
        }
        case 326: { /* '326' */
            return KnxDatapointMainType_DPT_4_1_1_BYTE_COMBINED_INFORMATION
        }
        case 327: { /* '327' */
            return KnxDatapointMainType_DPT_MBUS_ADDRESS
        }
        case 328: { /* '328' */
            return KnxDatapointMainType_DPT_3_BYTE_COLOUR_RGB
        }
        case 329: { /* '329' */
            return KnxDatapointMainType_DPT_LANGUAGE_CODE_ISO_639_1
        }
        case 33: { /* '33' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 330: { /* '330' */
            return KnxDatapointMainType_DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY
        }
        case 331: { /* '331' */
            return KnxDatapointMainType_DPT_PRIORITISED_MODE_CONTROL
        }
        case 332: { /* '332' */
            return KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_16_BIT
        }
        case 333: { /* '333' */
            return KnxDatapointMainType_DPT_CONFIGURATION_DIAGNOSTICS_8_BIT
        }
        case 334: { /* '334' */
            return KnxDatapointMainType_DPT_POSITIONS
        }
        case 335: { /* '335' */
            return KnxDatapointMainType_DPT_STATUS_32_BIT
        }
        case 336: { /* '336' */
            return KnxDatapointMainType_DPT_STATUS_48_BIT
        }
        case 337: { /* '337' */
            return KnxDatapointMainType_DPT_CONVERTER_STATUS
        }
        case 338: { /* '338' */
            return KnxDatapointMainType_DPT_CONVERTER_TEST_RESULT
        }
        case 339: { /* '339' */
            return KnxDatapointMainType_DPT_BATTERY_INFORMATION
        }
        case 34: { /* '34' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 340: { /* '340' */
            return KnxDatapointMainType_DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION
        }
        case 341: { /* '341' */
            return KnxDatapointMainType_DPT_STATUS_24_BIT
        }
        case 342: { /* '342' */
            return KnxDatapointMainType_DPT_COLOUR_RGBW
        }
        case 343: { /* '343' */
            return KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGBW
        }
        case 344: { /* '344' */
            return KnxDatapointMainType_DPT_RELATIVE_CONTROL_RGB
        }
        case 345: { /* '345' */
            return KnxDatapointMainType_DPT_F32F32
        }
        case 346: { /* '346' */
            return KnxDatapointMainType_DPT_F16F16F16F16
        }
        case 347: { /* '347' */
            return KnxDatapointMainType_DPT_F16F16F16F16
        }
        case 35: { /* '35' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 36: { /* '36' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 37: { /* '37' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 38: { /* '38' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 39: { /* '39' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 4: { /* '4' */
            return KnxDatapointMainType_DPT_32_BIT_SET
        }
        case 40: { /* '40' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 41: { /* '41' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 42: { /* '42' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 43: { /* '43' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 44: { /* '44' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 45: { /* '45' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 46: { /* '46' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 47: { /* '47' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 48: { /* '48' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 49: { /* '49' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 5: { /* '5' */
            return KnxDatapointMainType_DPT_64_BIT_SET
        }
        case 50: { /* '50' */
            return KnxDatapointMainType_DPT_1_BIT
        }
        case 51: { /* '51' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 52: { /* '52' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 53: { /* '53' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 54: { /* '54' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 55: { /* '55' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 56: { /* '56' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 57: { /* '57' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 58: { /* '58' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 59: { /* '59' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 6: { /* '6' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 60: { /* '60' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 61: { /* '61' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 62: { /* '62' */
            return KnxDatapointMainType_DPT_1_BIT_CONTROLLED
        }
        case 63: { /* '63' */
            return KnxDatapointMainType_DPT_3_BIT_CONTROLLED
        }
        case 64: { /* '64' */
            return KnxDatapointMainType_DPT_3_BIT_CONTROLLED
        }
        case 65: { /* '65' */
            return KnxDatapointMainType_DPT_CHARACTER
        }
        case 66: { /* '66' */
            return KnxDatapointMainType_DPT_CHARACTER
        }
        case 67: { /* '67' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 68: { /* '68' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 69: { /* '69' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 7: { /* '7' */
            return KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE
        }
        case 70: { /* '70' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 71: { /* '71' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 72: { /* '72' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 73: { /* '73' */
            return KnxDatapointMainType_DPT_8_BIT_UNSIGNED_VALUE
        }
        case 74: { /* '74' */
            return KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE
        }
        case 75: { /* '75' */
            return KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE
        }
        case 76: { /* '76' */
            return KnxDatapointMainType_DPT_8_BIT_SIGNED_VALUE
        }
        case 77: { /* '77' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 78: { /* '78' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 79: { /* '79' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 8: { /* '8' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 80: { /* '80' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 81: { /* '81' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 82: { /* '82' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 83: { /* '83' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 84: { /* '84' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 85: { /* '85' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 86: { /* '86' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 87: { /* '87' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 88: { /* '88' */
            return KnxDatapointMainType_DPT_2_BYTE_UNSIGNED_VALUE
        }
        case 89: { /* '89' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 9: { /* '9' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 90: { /* '90' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 91: { /* '91' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 92: { /* '92' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 93: { /* '93' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 94: { /* '94' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 95: { /* '95' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 96: { /* '96' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 97: { /* '97' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 98: { /* '98' */
            return KnxDatapointMainType_DPT_2_BYTE_SIGNED_VALUE
        }
        case 99: { /* '99' */
            return KnxDatapointMainType_DPT_2_BYTE_FLOAT_VALUE
        }
        default: {
            return 0
        }
    }
}
func KnxDatapointTypeByValue(value uint32) KnxDatapointType {
    switch value {
        case 0:
            return KnxDatapointType_DPT_UNKNOWN
        case 1:
            return KnxDatapointType_BOOL
        case 10:
            return KnxDatapointType_UDINT
        case 100:
            return KnxDatapointType_DPT_Value_Tempd
        case 101:
            return KnxDatapointType_DPT_Value_Tempa
        case 102:
            return KnxDatapointType_DPT_Value_Lux
        case 103:
            return KnxDatapointType_DPT_Value_Wsp
        case 104:
            return KnxDatapointType_DPT_Value_Pres
        case 105:
            return KnxDatapointType_DPT_Value_Humidity
        case 106:
            return KnxDatapointType_DPT_Value_AirQuality
        case 107:
            return KnxDatapointType_DPT_Value_AirFlow
        case 108:
            return KnxDatapointType_DPT_Value_Time1
        case 109:
            return KnxDatapointType_DPT_Value_Time2
        case 11:
            return KnxDatapointType_DINT
        case 110:
            return KnxDatapointType_DPT_Value_Volt
        case 111:
            return KnxDatapointType_DPT_Value_Curr
        case 112:
            return KnxDatapointType_DPT_PowerDensity
        case 113:
            return KnxDatapointType_DPT_KelvinPerPercent
        case 114:
            return KnxDatapointType_DPT_Power
        case 115:
            return KnxDatapointType_DPT_Value_Volume_Flow
        case 116:
            return KnxDatapointType_DPT_Rain_Amount
        case 117:
            return KnxDatapointType_DPT_Value_Temp_F
        case 118:
            return KnxDatapointType_DPT_Value_Wsp_kmh
        case 119:
            return KnxDatapointType_DPT_Value_Absolute_Humidity
        case 12:
            return KnxDatapointType_ULINT
        case 120:
            return KnxDatapointType_DPT_Concentration_ygm3
        case 121:
            return KnxDatapointType_DPT_TimeOfDay
        case 122:
            return KnxDatapointType_DPT_Date
        case 123:
            return KnxDatapointType_DPT_Value_4_Ucount
        case 124:
            return KnxDatapointType_DPT_LongTimePeriod_Sec
        case 125:
            return KnxDatapointType_DPT_LongTimePeriod_Min
        case 126:
            return KnxDatapointType_DPT_LongTimePeriod_Hrs
        case 127:
            return KnxDatapointType_DPT_VolumeLiquid_Litre
        case 128:
            return KnxDatapointType_DPT_Volume_m_3
        case 129:
            return KnxDatapointType_DPT_Value_4_Count
        case 13:
            return KnxDatapointType_LINT
        case 130:
            return KnxDatapointType_DPT_FlowRate_m3h
        case 131:
            return KnxDatapointType_DPT_ActiveEnergy
        case 132:
            return KnxDatapointType_DPT_ApparantEnergy
        case 133:
            return KnxDatapointType_DPT_ReactiveEnergy
        case 134:
            return KnxDatapointType_DPT_ActiveEnergy_kWh
        case 135:
            return KnxDatapointType_DPT_ApparantEnergy_kVAh
        case 136:
            return KnxDatapointType_DPT_ReactiveEnergy_kVARh
        case 137:
            return KnxDatapointType_DPT_ActiveEnergy_MWh
        case 138:
            return KnxDatapointType_DPT_LongDeltaTimeSec
        case 139:
            return KnxDatapointType_DPT_DeltaVolumeLiquid_Litre
        case 14:
            return KnxDatapointType_REAL
        case 140:
            return KnxDatapointType_DPT_DeltaVolume_m_3
        case 141:
            return KnxDatapointType_DPT_Value_Acceleration
        case 142:
            return KnxDatapointType_DPT_Value_Acceleration_Angular
        case 143:
            return KnxDatapointType_DPT_Value_Activation_Energy
        case 144:
            return KnxDatapointType_DPT_Value_Activity
        case 145:
            return KnxDatapointType_DPT_Value_Mol
        case 146:
            return KnxDatapointType_DPT_Value_Amplitude
        case 147:
            return KnxDatapointType_DPT_Value_AngleRad
        case 148:
            return KnxDatapointType_DPT_Value_AngleDeg
        case 149:
            return KnxDatapointType_DPT_Value_Angular_Momentum
        case 15:
            return KnxDatapointType_LREAL
        case 150:
            return KnxDatapointType_DPT_Value_Angular_Velocity
        case 151:
            return KnxDatapointType_DPT_Value_Area
        case 152:
            return KnxDatapointType_DPT_Value_Capacitance
        case 153:
            return KnxDatapointType_DPT_Value_Charge_DensitySurface
        case 154:
            return KnxDatapointType_DPT_Value_Charge_DensityVolume
        case 155:
            return KnxDatapointType_DPT_Value_Compressibility
        case 156:
            return KnxDatapointType_DPT_Value_Conductance
        case 157:
            return KnxDatapointType_DPT_Value_Electrical_Conductivity
        case 158:
            return KnxDatapointType_DPT_Value_Density
        case 159:
            return KnxDatapointType_DPT_Value_Electric_Charge
        case 16:
            return KnxDatapointType_CHAR
        case 160:
            return KnxDatapointType_DPT_Value_Electric_Current
        case 161:
            return KnxDatapointType_DPT_Value_Electric_CurrentDensity
        case 162:
            return KnxDatapointType_DPT_Value_Electric_DipoleMoment
        case 163:
            return KnxDatapointType_DPT_Value_Electric_Displacement
        case 164:
            return KnxDatapointType_DPT_Value_Electric_FieldStrength
        case 165:
            return KnxDatapointType_DPT_Value_Electric_Flux
        case 166:
            return KnxDatapointType_DPT_Value_Electric_FluxDensity
        case 167:
            return KnxDatapointType_DPT_Value_Electric_Polarization
        case 168:
            return KnxDatapointType_DPT_Value_Electric_Potential
        case 169:
            return KnxDatapointType_DPT_Value_Electric_PotentialDifference
        case 17:
            return KnxDatapointType_WCHAR
        case 170:
            return KnxDatapointType_DPT_Value_ElectromagneticMoment
        case 171:
            return KnxDatapointType_DPT_Value_Electromotive_Force
        case 172:
            return KnxDatapointType_DPT_Value_Energy
        case 173:
            return KnxDatapointType_DPT_Value_Force
        case 174:
            return KnxDatapointType_DPT_Value_Frequency
        case 175:
            return KnxDatapointType_DPT_Value_Angular_Frequency
        case 176:
            return KnxDatapointType_DPT_Value_Heat_Capacity
        case 177:
            return KnxDatapointType_DPT_Value_Heat_FlowRate
        case 178:
            return KnxDatapointType_DPT_Value_Heat_Quantity
        case 179:
            return KnxDatapointType_DPT_Value_Impedance
        case 18:
            return KnxDatapointType_STRING
        case 180:
            return KnxDatapointType_DPT_Value_Length
        case 181:
            return KnxDatapointType_DPT_Value_Light_Quantity
        case 182:
            return KnxDatapointType_DPT_Value_Luminance
        case 183:
            return KnxDatapointType_DPT_Value_Luminous_Flux
        case 184:
            return KnxDatapointType_DPT_Value_Luminous_Intensity
        case 185:
            return KnxDatapointType_DPT_Value_Magnetic_FieldStrength
        case 186:
            return KnxDatapointType_DPT_Value_Magnetic_Flux
        case 187:
            return KnxDatapointType_DPT_Value_Magnetic_FluxDensity
        case 188:
            return KnxDatapointType_DPT_Value_Magnetic_Moment
        case 189:
            return KnxDatapointType_DPT_Value_Magnetic_Polarization
        case 19:
            return KnxDatapointType_WSTRING
        case 190:
            return KnxDatapointType_DPT_Value_Magnetization
        case 191:
            return KnxDatapointType_DPT_Value_MagnetomotiveForce
        case 192:
            return KnxDatapointType_DPT_Value_Mass
        case 193:
            return KnxDatapointType_DPT_Value_MassFlux
        case 194:
            return KnxDatapointType_DPT_Value_Momentum
        case 195:
            return KnxDatapointType_DPT_Value_Phase_AngleRad
        case 196:
            return KnxDatapointType_DPT_Value_Phase_AngleDeg
        case 197:
            return KnxDatapointType_DPT_Value_Power
        case 198:
            return KnxDatapointType_DPT_Value_Power_Factor
        case 199:
            return KnxDatapointType_DPT_Value_Pressure
        case 2:
            return KnxDatapointType_BYTE
        case 20:
            return KnxDatapointType_TIME
        case 200:
            return KnxDatapointType_DPT_Value_Reactance
        case 201:
            return KnxDatapointType_DPT_Value_Resistance
        case 202:
            return KnxDatapointType_DPT_Value_Resistivity
        case 203:
            return KnxDatapointType_DPT_Value_SelfInductance
        case 204:
            return KnxDatapointType_DPT_Value_SolidAngle
        case 205:
            return KnxDatapointType_DPT_Value_Sound_Intensity
        case 206:
            return KnxDatapointType_DPT_Value_Speed
        case 207:
            return KnxDatapointType_DPT_Value_Stress
        case 208:
            return KnxDatapointType_DPT_Value_Surface_Tension
        case 209:
            return KnxDatapointType_DPT_Value_Common_Temperature
        case 21:
            return KnxDatapointType_LTIME
        case 210:
            return KnxDatapointType_DPT_Value_Absolute_Temperature
        case 211:
            return KnxDatapointType_DPT_Value_TemperatureDifference
        case 212:
            return KnxDatapointType_DPT_Value_Thermal_Capacity
        case 213:
            return KnxDatapointType_DPT_Value_Thermal_Conductivity
        case 214:
            return KnxDatapointType_DPT_Value_ThermoelectricPower
        case 215:
            return KnxDatapointType_DPT_Value_Time
        case 216:
            return KnxDatapointType_DPT_Value_Torque
        case 217:
            return KnxDatapointType_DPT_Value_Volume
        case 218:
            return KnxDatapointType_DPT_Value_Volume_Flux
        case 219:
            return KnxDatapointType_DPT_Value_Weight
        case 22:
            return KnxDatapointType_DATE
        case 220:
            return KnxDatapointType_DPT_Value_Work
        case 221:
            return KnxDatapointType_DPT_Volume_Flux_Meter
        case 222:
            return KnxDatapointType_DPT_Volume_Flux_ls
        case 223:
            return KnxDatapointType_DPT_Access_Data
        case 224:
            return KnxDatapointType_DPT_String_ASCII
        case 225:
            return KnxDatapointType_DPT_String_8859_1
        case 226:
            return KnxDatapointType_DPT_SceneNumber
        case 227:
            return KnxDatapointType_DPT_SceneControl
        case 228:
            return KnxDatapointType_DPT_DateTime
        case 229:
            return KnxDatapointType_DPT_SCLOMode
        case 23:
            return KnxDatapointType_TIME_OF_DAY
        case 230:
            return KnxDatapointType_DPT_BuildingMode
        case 231:
            return KnxDatapointType_DPT_OccMode
        case 232:
            return KnxDatapointType_DPT_Priority
        case 233:
            return KnxDatapointType_DPT_LightApplicationMode
        case 234:
            return KnxDatapointType_DPT_ApplicationArea
        case 235:
            return KnxDatapointType_DPT_AlarmClassType
        case 236:
            return KnxDatapointType_DPT_PSUMode
        case 237:
            return KnxDatapointType_DPT_ErrorClass_System
        case 238:
            return KnxDatapointType_DPT_ErrorClass_HVAC
        case 239:
            return KnxDatapointType_DPT_Time_Delay
        case 24:
            return KnxDatapointType_TOD
        case 240:
            return KnxDatapointType_DPT_Beaufort_Wind_Force_Scale
        case 241:
            return KnxDatapointType_DPT_SensorSelect
        case 242:
            return KnxDatapointType_DPT_ActuatorConnectType
        case 243:
            return KnxDatapointType_DPT_Cloud_Cover
        case 244:
            return KnxDatapointType_DPT_PowerReturnMode
        case 245:
            return KnxDatapointType_DPT_FuelType
        case 246:
            return KnxDatapointType_DPT_BurnerType
        case 247:
            return KnxDatapointType_DPT_HVACMode
        case 248:
            return KnxDatapointType_DPT_DHWMode
        case 249:
            return KnxDatapointType_DPT_LoadPriority
        case 25:
            return KnxDatapointType_DATE_AND_TIME
        case 250:
            return KnxDatapointType_DPT_HVACContrMode
        case 251:
            return KnxDatapointType_DPT_HVACEmergMode
        case 252:
            return KnxDatapointType_DPT_ChangeoverMode
        case 253:
            return KnxDatapointType_DPT_ValveMode
        case 254:
            return KnxDatapointType_DPT_DamperMode
        case 255:
            return KnxDatapointType_DPT_HeaterMode
        case 256:
            return KnxDatapointType_DPT_FanMode
        case 257:
            return KnxDatapointType_DPT_MasterSlaveMode
        case 258:
            return KnxDatapointType_DPT_StatusRoomSetp
        case 259:
            return KnxDatapointType_DPT_Metering_DeviceType
        case 26:
            return KnxDatapointType_DT
        case 260:
            return KnxDatapointType_DPT_HumDehumMode
        case 261:
            return KnxDatapointType_DPT_EnableHCStage
        case 262:
            return KnxDatapointType_DPT_ADAType
        case 263:
            return KnxDatapointType_DPT_BackupMode
        case 264:
            return KnxDatapointType_DPT_StartSynchronization
        case 265:
            return KnxDatapointType_DPT_Behaviour_Lock_Unlock
        case 266:
            return KnxDatapointType_DPT_Behaviour_Bus_Power_Up_Down
        case 267:
            return KnxDatapointType_DPT_DALI_Fade_Time
        case 268:
            return KnxDatapointType_DPT_BlinkingMode
        case 269:
            return KnxDatapointType_DPT_LightControlMode
        case 27:
            return KnxDatapointType_DPT_Switch
        case 270:
            return KnxDatapointType_DPT_SwitchPBModel
        case 271:
            return KnxDatapointType_DPT_PBAction
        case 272:
            return KnxDatapointType_DPT_DimmPBModel
        case 273:
            return KnxDatapointType_DPT_SwitchOnMode
        case 274:
            return KnxDatapointType_DPT_LoadTypeSet
        case 275:
            return KnxDatapointType_DPT_LoadTypeDetected
        case 276:
            return KnxDatapointType_DPT_Converter_Test_Control
        case 277:
            return KnxDatapointType_DPT_SABExcept_Behaviour
        case 278:
            return KnxDatapointType_DPT_SABBehaviour_Lock_Unlock
        case 279:
            return KnxDatapointType_DPT_SSSBMode
        case 28:
            return KnxDatapointType_DPT_Bool
        case 280:
            return KnxDatapointType_DPT_BlindsControlMode
        case 281:
            return KnxDatapointType_DPT_CommMode
        case 282:
            return KnxDatapointType_DPT_AddInfoTypes
        case 283:
            return KnxDatapointType_DPT_RF_ModeSelect
        case 284:
            return KnxDatapointType_DPT_RF_FilterSelect
        case 285:
            return KnxDatapointType_DPT_StatusGen
        case 286:
            return KnxDatapointType_DPT_Device_Control
        case 287:
            return KnxDatapointType_DPT_ForceSign
        case 288:
            return KnxDatapointType_DPT_ForceSignCool
        case 289:
            return KnxDatapointType_DPT_StatusRHC
        case 29:
            return KnxDatapointType_DPT_Enable
        case 290:
            return KnxDatapointType_DPT_StatusSDHWC
        case 291:
            return KnxDatapointType_DPT_FuelTypeSet
        case 292:
            return KnxDatapointType_DPT_StatusRCC
        case 293:
            return KnxDatapointType_DPT_StatusAHU
        case 294:
            return KnxDatapointType_DPT_CombinedStatus_RTSM
        case 295:
            return KnxDatapointType_DPT_LightActuatorErrorInfo
        case 296:
            return KnxDatapointType_DPT_RF_ModeInfo
        case 297:
            return KnxDatapointType_DPT_RF_FilterInfo
        case 298:
            return KnxDatapointType_DPT_Channel_Activation_8
        case 299:
            return KnxDatapointType_DPT_StatusDHWC
        case 3:
            return KnxDatapointType_WORD
        case 30:
            return KnxDatapointType_DPT_Ramp
        case 300:
            return KnxDatapointType_DPT_StatusRHCC
        case 301:
            return KnxDatapointType_DPT_CombinedStatus_HVA
        case 302:
            return KnxDatapointType_DPT_CombinedStatus_RTC
        case 303:
            return KnxDatapointType_DPT_Media
        case 304:
            return KnxDatapointType_DPT_Channel_Activation_16
        case 305:
            return KnxDatapointType_DPT_OnOffAction
        case 306:
            return KnxDatapointType_DPT_Alarm_Reaction
        case 307:
            return KnxDatapointType_DPT_UpDown_Action
        case 308:
            return KnxDatapointType_DPT_HVAC_PB_Action
        case 309:
            return KnxDatapointType_DPT_DoubleNibble
        case 31:
            return KnxDatapointType_DPT_Alarm
        case 310:
            return KnxDatapointType_DPT_SceneInfo
        case 311:
            return KnxDatapointType_DPT_CombinedInfoOnOff
        case 312:
            return KnxDatapointType_DPT_ActiveEnergy_V64
        case 313:
            return KnxDatapointType_DPT_ApparantEnergy_V64
        case 314:
            return KnxDatapointType_DPT_ReactiveEnergy_V64
        case 315:
            return KnxDatapointType_DPT_Channel_Activation_24
        case 316:
            return KnxDatapointType_DPT_HVACModeNext
        case 317:
            return KnxDatapointType_DPT_DHWModeNext
        case 318:
            return KnxDatapointType_DPT_OccModeNext
        case 319:
            return KnxDatapointType_DPT_BuildingModeNext
        case 32:
            return KnxDatapointType_DPT_BinaryValue
        case 320:
            return KnxDatapointType_DPT_Version
        case 321:
            return KnxDatapointType_DPT_AlarmInfo
        case 322:
            return KnxDatapointType_DPT_TempRoomSetpSetF16_3
        case 323:
            return KnxDatapointType_DPT_TempRoomSetpSetShiftF16_3
        case 324:
            return KnxDatapointType_DPT_Scaling_Speed
        case 325:
            return KnxDatapointType_DPT_Scaling_Step_Time
        case 326:
            return KnxDatapointType_DPT_MeteringValue
        case 327:
            return KnxDatapointType_DPT_MBus_Address
        case 328:
            return KnxDatapointType_DPT_Colour_RGB
        case 329:
            return KnxDatapointType_DPT_LanguageCodeAlpha2_ASCII
        case 33:
            return KnxDatapointType_DPT_Step
        case 330:
            return KnxDatapointType_DPT_Tariff_ActiveEnergy
        case 331:
            return KnxDatapointType_DPT_Prioritised_Mode_Control
        case 332:
            return KnxDatapointType_DPT_DALI_Control_Gear_Diagnostic
        case 333:
            return KnxDatapointType_DPT_DALI_Diagnostics
        case 334:
            return KnxDatapointType_DPT_CombinedPosition
        case 335:
            return KnxDatapointType_DPT_StatusSAB
        case 336:
            return KnxDatapointType_DPT_Colour_xyY
        case 337:
            return KnxDatapointType_DPT_Converter_Status
        case 338:
            return KnxDatapointType_DPT_Converter_Test_Result
        case 339:
            return KnxDatapointType_DPT_Battery_Info
        case 34:
            return KnxDatapointType_DPT_UpDown
        case 340:
            return KnxDatapointType_DPT_Brightness_Colour_Temperature_Transition
        case 341:
            return KnxDatapointType_DPT_Brightness_Colour_Temperature_Control
        case 342:
            return KnxDatapointType_DPT_Colour_RGBW
        case 343:
            return KnxDatapointType_DPT_Relative_Control_RGBW
        case 344:
            return KnxDatapointType_DPT_Relative_Control_RGB
        case 345:
            return KnxDatapointType_DPT_GeographicalLocation
        case 346:
            return KnxDatapointType_DPT_TempRoomSetpSetF16_4
        case 347:
            return KnxDatapointType_DPT_TempRoomSetpSetShiftF16_4
        case 35:
            return KnxDatapointType_DPT_OpenClose
        case 36:
            return KnxDatapointType_DPT_Start
        case 37:
            return KnxDatapointType_DPT_State
        case 38:
            return KnxDatapointType_DPT_Invert
        case 39:
            return KnxDatapointType_DPT_DimSendStyle
        case 4:
            return KnxDatapointType_DWORD
        case 40:
            return KnxDatapointType_DPT_InputSource
        case 41:
            return KnxDatapointType_DPT_Reset
        case 42:
            return KnxDatapointType_DPT_Ack
        case 43:
            return KnxDatapointType_DPT_Trigger
        case 44:
            return KnxDatapointType_DPT_Occupancy
        case 45:
            return KnxDatapointType_DPT_Window_Door
        case 46:
            return KnxDatapointType_DPT_LogicalFunction
        case 47:
            return KnxDatapointType_DPT_Scene_AB
        case 48:
            return KnxDatapointType_DPT_ShutterBlinds_Mode
        case 49:
            return KnxDatapointType_DPT_DayNight
        case 5:
            return KnxDatapointType_LWORD
        case 50:
            return KnxDatapointType_DPT_Heat_Cool
        case 51:
            return KnxDatapointType_DPT_Switch_Control
        case 52:
            return KnxDatapointType_DPT_Bool_Control
        case 53:
            return KnxDatapointType_DPT_Enable_Control
        case 54:
            return KnxDatapointType_DPT_Ramp_Control
        case 55:
            return KnxDatapointType_DPT_Alarm_Control
        case 56:
            return KnxDatapointType_DPT_BinaryValue_Control
        case 57:
            return KnxDatapointType_DPT_Step_Control
        case 58:
            return KnxDatapointType_DPT_Direction1_Control
        case 59:
            return KnxDatapointType_DPT_Direction2_Control
        case 6:
            return KnxDatapointType_USINT
        case 60:
            return KnxDatapointType_DPT_Start_Control
        case 61:
            return KnxDatapointType_DPT_State_Control
        case 62:
            return KnxDatapointType_DPT_Invert_Control
        case 63:
            return KnxDatapointType_DPT_Control_Dimming
        case 64:
            return KnxDatapointType_DPT_Control_Blinds
        case 65:
            return KnxDatapointType_DPT_Char_ASCII
        case 66:
            return KnxDatapointType_DPT_Char_8859_1
        case 67:
            return KnxDatapointType_DPT_Scaling
        case 68:
            return KnxDatapointType_DPT_Angle
        case 69:
            return KnxDatapointType_DPT_Percent_U8
        case 7:
            return KnxDatapointType_SINT
        case 70:
            return KnxDatapointType_DPT_DecimalFactor
        case 71:
            return KnxDatapointType_DPT_Tariff
        case 72:
            return KnxDatapointType_DPT_Value_1_Ucount
        case 73:
            return KnxDatapointType_DPT_FanStage
        case 74:
            return KnxDatapointType_DPT_Percent_V8
        case 75:
            return KnxDatapointType_DPT_Value_1_Count
        case 76:
            return KnxDatapointType_DPT_Status_Mode3
        case 77:
            return KnxDatapointType_DPT_Value_2_Ucount
        case 78:
            return KnxDatapointType_DPT_TimePeriodMsec
        case 79:
            return KnxDatapointType_DPT_TimePeriod10Msec
        case 8:
            return KnxDatapointType_UINT
        case 80:
            return KnxDatapointType_DPT_TimePeriod100Msec
        case 81:
            return KnxDatapointType_DPT_TimePeriodSec
        case 82:
            return KnxDatapointType_DPT_TimePeriodMin
        case 83:
            return KnxDatapointType_DPT_TimePeriodHrs
        case 84:
            return KnxDatapointType_DPT_PropDataType
        case 85:
            return KnxDatapointType_DPT_Length_mm
        case 86:
            return KnxDatapointType_DPT_UElCurrentmA
        case 87:
            return KnxDatapointType_DPT_Brightness
        case 88:
            return KnxDatapointType_DPT_Absolute_Colour_Temperature
        case 89:
            return KnxDatapointType_DPT_Value_2_Count
        case 9:
            return KnxDatapointType_INT
        case 90:
            return KnxDatapointType_DPT_DeltaTimeMsec
        case 91:
            return KnxDatapointType_DPT_DeltaTime10Msec
        case 92:
            return KnxDatapointType_DPT_DeltaTime100Msec
        case 93:
            return KnxDatapointType_DPT_DeltaTimeSec
        case 94:
            return KnxDatapointType_DPT_DeltaTimeMin
        case 95:
            return KnxDatapointType_DPT_DeltaTimeHrs
        case 96:
            return KnxDatapointType_DPT_Percent_V16
        case 97:
            return KnxDatapointType_DPT_Rotation_Angle
        case 98:
            return KnxDatapointType_DPT_Length_m
        case 99:
            return KnxDatapointType_DPT_Value_Temp
    }
    return 0
}

func KnxDatapointTypeByName(value string) KnxDatapointType {
    switch value {
    case "DPT_UNKNOWN":
        return KnxDatapointType_DPT_UNKNOWN
    case "BOOL":
        return KnxDatapointType_BOOL
    case "UDINT":
        return KnxDatapointType_UDINT
    case "DPT_Value_Tempd":
        return KnxDatapointType_DPT_Value_Tempd
    case "DPT_Value_Tempa":
        return KnxDatapointType_DPT_Value_Tempa
    case "DPT_Value_Lux":
        return KnxDatapointType_DPT_Value_Lux
    case "DPT_Value_Wsp":
        return KnxDatapointType_DPT_Value_Wsp
    case "DPT_Value_Pres":
        return KnxDatapointType_DPT_Value_Pres
    case "DPT_Value_Humidity":
        return KnxDatapointType_DPT_Value_Humidity
    case "DPT_Value_AirQuality":
        return KnxDatapointType_DPT_Value_AirQuality
    case "DPT_Value_AirFlow":
        return KnxDatapointType_DPT_Value_AirFlow
    case "DPT_Value_Time1":
        return KnxDatapointType_DPT_Value_Time1
    case "DPT_Value_Time2":
        return KnxDatapointType_DPT_Value_Time2
    case "DINT":
        return KnxDatapointType_DINT
    case "DPT_Value_Volt":
        return KnxDatapointType_DPT_Value_Volt
    case "DPT_Value_Curr":
        return KnxDatapointType_DPT_Value_Curr
    case "DPT_PowerDensity":
        return KnxDatapointType_DPT_PowerDensity
    case "DPT_KelvinPerPercent":
        return KnxDatapointType_DPT_KelvinPerPercent
    case "DPT_Power":
        return KnxDatapointType_DPT_Power
    case "DPT_Value_Volume_Flow":
        return KnxDatapointType_DPT_Value_Volume_Flow
    case "DPT_Rain_Amount":
        return KnxDatapointType_DPT_Rain_Amount
    case "DPT_Value_Temp_F":
        return KnxDatapointType_DPT_Value_Temp_F
    case "DPT_Value_Wsp_kmh":
        return KnxDatapointType_DPT_Value_Wsp_kmh
    case "DPT_Value_Absolute_Humidity":
        return KnxDatapointType_DPT_Value_Absolute_Humidity
    case "ULINT":
        return KnxDatapointType_ULINT
    case "DPT_Concentration_ygm3":
        return KnxDatapointType_DPT_Concentration_ygm3
    case "DPT_TimeOfDay":
        return KnxDatapointType_DPT_TimeOfDay
    case "DPT_Date":
        return KnxDatapointType_DPT_Date
    case "DPT_Value_4_Ucount":
        return KnxDatapointType_DPT_Value_4_Ucount
    case "DPT_LongTimePeriod_Sec":
        return KnxDatapointType_DPT_LongTimePeriod_Sec
    case "DPT_LongTimePeriod_Min":
        return KnxDatapointType_DPT_LongTimePeriod_Min
    case "DPT_LongTimePeriod_Hrs":
        return KnxDatapointType_DPT_LongTimePeriod_Hrs
    case "DPT_VolumeLiquid_Litre":
        return KnxDatapointType_DPT_VolumeLiquid_Litre
    case "DPT_Volume_m_3":
        return KnxDatapointType_DPT_Volume_m_3
    case "DPT_Value_4_Count":
        return KnxDatapointType_DPT_Value_4_Count
    case "LINT":
        return KnxDatapointType_LINT
    case "DPT_FlowRate_m3h":
        return KnxDatapointType_DPT_FlowRate_m3h
    case "DPT_ActiveEnergy":
        return KnxDatapointType_DPT_ActiveEnergy
    case "DPT_ApparantEnergy":
        return KnxDatapointType_DPT_ApparantEnergy
    case "DPT_ReactiveEnergy":
        return KnxDatapointType_DPT_ReactiveEnergy
    case "DPT_ActiveEnergy_kWh":
        return KnxDatapointType_DPT_ActiveEnergy_kWh
    case "DPT_ApparantEnergy_kVAh":
        return KnxDatapointType_DPT_ApparantEnergy_kVAh
    case "DPT_ReactiveEnergy_kVARh":
        return KnxDatapointType_DPT_ReactiveEnergy_kVARh
    case "DPT_ActiveEnergy_MWh":
        return KnxDatapointType_DPT_ActiveEnergy_MWh
    case "DPT_LongDeltaTimeSec":
        return KnxDatapointType_DPT_LongDeltaTimeSec
    case "DPT_DeltaVolumeLiquid_Litre":
        return KnxDatapointType_DPT_DeltaVolumeLiquid_Litre
    case "REAL":
        return KnxDatapointType_REAL
    case "DPT_DeltaVolume_m_3":
        return KnxDatapointType_DPT_DeltaVolume_m_3
    case "DPT_Value_Acceleration":
        return KnxDatapointType_DPT_Value_Acceleration
    case "DPT_Value_Acceleration_Angular":
        return KnxDatapointType_DPT_Value_Acceleration_Angular
    case "DPT_Value_Activation_Energy":
        return KnxDatapointType_DPT_Value_Activation_Energy
    case "DPT_Value_Activity":
        return KnxDatapointType_DPT_Value_Activity
    case "DPT_Value_Mol":
        return KnxDatapointType_DPT_Value_Mol
    case "DPT_Value_Amplitude":
        return KnxDatapointType_DPT_Value_Amplitude
    case "DPT_Value_AngleRad":
        return KnxDatapointType_DPT_Value_AngleRad
    case "DPT_Value_AngleDeg":
        return KnxDatapointType_DPT_Value_AngleDeg
    case "DPT_Value_Angular_Momentum":
        return KnxDatapointType_DPT_Value_Angular_Momentum
    case "LREAL":
        return KnxDatapointType_LREAL
    case "DPT_Value_Angular_Velocity":
        return KnxDatapointType_DPT_Value_Angular_Velocity
    case "DPT_Value_Area":
        return KnxDatapointType_DPT_Value_Area
    case "DPT_Value_Capacitance":
        return KnxDatapointType_DPT_Value_Capacitance
    case "DPT_Value_Charge_DensitySurface":
        return KnxDatapointType_DPT_Value_Charge_DensitySurface
    case "DPT_Value_Charge_DensityVolume":
        return KnxDatapointType_DPT_Value_Charge_DensityVolume
    case "DPT_Value_Compressibility":
        return KnxDatapointType_DPT_Value_Compressibility
    case "DPT_Value_Conductance":
        return KnxDatapointType_DPT_Value_Conductance
    case "DPT_Value_Electrical_Conductivity":
        return KnxDatapointType_DPT_Value_Electrical_Conductivity
    case "DPT_Value_Density":
        return KnxDatapointType_DPT_Value_Density
    case "DPT_Value_Electric_Charge":
        return KnxDatapointType_DPT_Value_Electric_Charge
    case "CHAR":
        return KnxDatapointType_CHAR
    case "DPT_Value_Electric_Current":
        return KnxDatapointType_DPT_Value_Electric_Current
    case "DPT_Value_Electric_CurrentDensity":
        return KnxDatapointType_DPT_Value_Electric_CurrentDensity
    case "DPT_Value_Electric_DipoleMoment":
        return KnxDatapointType_DPT_Value_Electric_DipoleMoment
    case "DPT_Value_Electric_Displacement":
        return KnxDatapointType_DPT_Value_Electric_Displacement
    case "DPT_Value_Electric_FieldStrength":
        return KnxDatapointType_DPT_Value_Electric_FieldStrength
    case "DPT_Value_Electric_Flux":
        return KnxDatapointType_DPT_Value_Electric_Flux
    case "DPT_Value_Electric_FluxDensity":
        return KnxDatapointType_DPT_Value_Electric_FluxDensity
    case "DPT_Value_Electric_Polarization":
        return KnxDatapointType_DPT_Value_Electric_Polarization
    case "DPT_Value_Electric_Potential":
        return KnxDatapointType_DPT_Value_Electric_Potential
    case "DPT_Value_Electric_PotentialDifference":
        return KnxDatapointType_DPT_Value_Electric_PotentialDifference
    case "WCHAR":
        return KnxDatapointType_WCHAR
    case "DPT_Value_ElectromagneticMoment":
        return KnxDatapointType_DPT_Value_ElectromagneticMoment
    case "DPT_Value_Electromotive_Force":
        return KnxDatapointType_DPT_Value_Electromotive_Force
    case "DPT_Value_Energy":
        return KnxDatapointType_DPT_Value_Energy
    case "DPT_Value_Force":
        return KnxDatapointType_DPT_Value_Force
    case "DPT_Value_Frequency":
        return KnxDatapointType_DPT_Value_Frequency
    case "DPT_Value_Angular_Frequency":
        return KnxDatapointType_DPT_Value_Angular_Frequency
    case "DPT_Value_Heat_Capacity":
        return KnxDatapointType_DPT_Value_Heat_Capacity
    case "DPT_Value_Heat_FlowRate":
        return KnxDatapointType_DPT_Value_Heat_FlowRate
    case "DPT_Value_Heat_Quantity":
        return KnxDatapointType_DPT_Value_Heat_Quantity
    case "DPT_Value_Impedance":
        return KnxDatapointType_DPT_Value_Impedance
    case "STRING":
        return KnxDatapointType_STRING
    case "DPT_Value_Length":
        return KnxDatapointType_DPT_Value_Length
    case "DPT_Value_Light_Quantity":
        return KnxDatapointType_DPT_Value_Light_Quantity
    case "DPT_Value_Luminance":
        return KnxDatapointType_DPT_Value_Luminance
    case "DPT_Value_Luminous_Flux":
        return KnxDatapointType_DPT_Value_Luminous_Flux
    case "DPT_Value_Luminous_Intensity":
        return KnxDatapointType_DPT_Value_Luminous_Intensity
    case "DPT_Value_Magnetic_FieldStrength":
        return KnxDatapointType_DPT_Value_Magnetic_FieldStrength
    case "DPT_Value_Magnetic_Flux":
        return KnxDatapointType_DPT_Value_Magnetic_Flux
    case "DPT_Value_Magnetic_FluxDensity":
        return KnxDatapointType_DPT_Value_Magnetic_FluxDensity
    case "DPT_Value_Magnetic_Moment":
        return KnxDatapointType_DPT_Value_Magnetic_Moment
    case "DPT_Value_Magnetic_Polarization":
        return KnxDatapointType_DPT_Value_Magnetic_Polarization
    case "WSTRING":
        return KnxDatapointType_WSTRING
    case "DPT_Value_Magnetization":
        return KnxDatapointType_DPT_Value_Magnetization
    case "DPT_Value_MagnetomotiveForce":
        return KnxDatapointType_DPT_Value_MagnetomotiveForce
    case "DPT_Value_Mass":
        return KnxDatapointType_DPT_Value_Mass
    case "DPT_Value_MassFlux":
        return KnxDatapointType_DPT_Value_MassFlux
    case "DPT_Value_Momentum":
        return KnxDatapointType_DPT_Value_Momentum
    case "DPT_Value_Phase_AngleRad":
        return KnxDatapointType_DPT_Value_Phase_AngleRad
    case "DPT_Value_Phase_AngleDeg":
        return KnxDatapointType_DPT_Value_Phase_AngleDeg
    case "DPT_Value_Power":
        return KnxDatapointType_DPT_Value_Power
    case "DPT_Value_Power_Factor":
        return KnxDatapointType_DPT_Value_Power_Factor
    case "DPT_Value_Pressure":
        return KnxDatapointType_DPT_Value_Pressure
    case "BYTE":
        return KnxDatapointType_BYTE
    case "TIME":
        return KnxDatapointType_TIME
    case "DPT_Value_Reactance":
        return KnxDatapointType_DPT_Value_Reactance
    case "DPT_Value_Resistance":
        return KnxDatapointType_DPT_Value_Resistance
    case "DPT_Value_Resistivity":
        return KnxDatapointType_DPT_Value_Resistivity
    case "DPT_Value_SelfInductance":
        return KnxDatapointType_DPT_Value_SelfInductance
    case "DPT_Value_SolidAngle":
        return KnxDatapointType_DPT_Value_SolidAngle
    case "DPT_Value_Sound_Intensity":
        return KnxDatapointType_DPT_Value_Sound_Intensity
    case "DPT_Value_Speed":
        return KnxDatapointType_DPT_Value_Speed
    case "DPT_Value_Stress":
        return KnxDatapointType_DPT_Value_Stress
    case "DPT_Value_Surface_Tension":
        return KnxDatapointType_DPT_Value_Surface_Tension
    case "DPT_Value_Common_Temperature":
        return KnxDatapointType_DPT_Value_Common_Temperature
    case "LTIME":
        return KnxDatapointType_LTIME
    case "DPT_Value_Absolute_Temperature":
        return KnxDatapointType_DPT_Value_Absolute_Temperature
    case "DPT_Value_TemperatureDifference":
        return KnxDatapointType_DPT_Value_TemperatureDifference
    case "DPT_Value_Thermal_Capacity":
        return KnxDatapointType_DPT_Value_Thermal_Capacity
    case "DPT_Value_Thermal_Conductivity":
        return KnxDatapointType_DPT_Value_Thermal_Conductivity
    case "DPT_Value_ThermoelectricPower":
        return KnxDatapointType_DPT_Value_ThermoelectricPower
    case "DPT_Value_Time":
        return KnxDatapointType_DPT_Value_Time
    case "DPT_Value_Torque":
        return KnxDatapointType_DPT_Value_Torque
    case "DPT_Value_Volume":
        return KnxDatapointType_DPT_Value_Volume
    case "DPT_Value_Volume_Flux":
        return KnxDatapointType_DPT_Value_Volume_Flux
    case "DPT_Value_Weight":
        return KnxDatapointType_DPT_Value_Weight
    case "DATE":
        return KnxDatapointType_DATE
    case "DPT_Value_Work":
        return KnxDatapointType_DPT_Value_Work
    case "DPT_Volume_Flux_Meter":
        return KnxDatapointType_DPT_Volume_Flux_Meter
    case "DPT_Volume_Flux_ls":
        return KnxDatapointType_DPT_Volume_Flux_ls
    case "DPT_Access_Data":
        return KnxDatapointType_DPT_Access_Data
    case "DPT_String_ASCII":
        return KnxDatapointType_DPT_String_ASCII
    case "DPT_String_8859_1":
        return KnxDatapointType_DPT_String_8859_1
    case "DPT_SceneNumber":
        return KnxDatapointType_DPT_SceneNumber
    case "DPT_SceneControl":
        return KnxDatapointType_DPT_SceneControl
    case "DPT_DateTime":
        return KnxDatapointType_DPT_DateTime
    case "DPT_SCLOMode":
        return KnxDatapointType_DPT_SCLOMode
    case "TIME_OF_DAY":
        return KnxDatapointType_TIME_OF_DAY
    case "DPT_BuildingMode":
        return KnxDatapointType_DPT_BuildingMode
    case "DPT_OccMode":
        return KnxDatapointType_DPT_OccMode
    case "DPT_Priority":
        return KnxDatapointType_DPT_Priority
    case "DPT_LightApplicationMode":
        return KnxDatapointType_DPT_LightApplicationMode
    case "DPT_ApplicationArea":
        return KnxDatapointType_DPT_ApplicationArea
    case "DPT_AlarmClassType":
        return KnxDatapointType_DPT_AlarmClassType
    case "DPT_PSUMode":
        return KnxDatapointType_DPT_PSUMode
    case "DPT_ErrorClass_System":
        return KnxDatapointType_DPT_ErrorClass_System
    case "DPT_ErrorClass_HVAC":
        return KnxDatapointType_DPT_ErrorClass_HVAC
    case "DPT_Time_Delay":
        return KnxDatapointType_DPT_Time_Delay
    case "TOD":
        return KnxDatapointType_TOD
    case "DPT_Beaufort_Wind_Force_Scale":
        return KnxDatapointType_DPT_Beaufort_Wind_Force_Scale
    case "DPT_SensorSelect":
        return KnxDatapointType_DPT_SensorSelect
    case "DPT_ActuatorConnectType":
        return KnxDatapointType_DPT_ActuatorConnectType
    case "DPT_Cloud_Cover":
        return KnxDatapointType_DPT_Cloud_Cover
    case "DPT_PowerReturnMode":
        return KnxDatapointType_DPT_PowerReturnMode
    case "DPT_FuelType":
        return KnxDatapointType_DPT_FuelType
    case "DPT_BurnerType":
        return KnxDatapointType_DPT_BurnerType
    case "DPT_HVACMode":
        return KnxDatapointType_DPT_HVACMode
    case "DPT_DHWMode":
        return KnxDatapointType_DPT_DHWMode
    case "DPT_LoadPriority":
        return KnxDatapointType_DPT_LoadPriority
    case "DATE_AND_TIME":
        return KnxDatapointType_DATE_AND_TIME
    case "DPT_HVACContrMode":
        return KnxDatapointType_DPT_HVACContrMode
    case "DPT_HVACEmergMode":
        return KnxDatapointType_DPT_HVACEmergMode
    case "DPT_ChangeoverMode":
        return KnxDatapointType_DPT_ChangeoverMode
    case "DPT_ValveMode":
        return KnxDatapointType_DPT_ValveMode
    case "DPT_DamperMode":
        return KnxDatapointType_DPT_DamperMode
    case "DPT_HeaterMode":
        return KnxDatapointType_DPT_HeaterMode
    case "DPT_FanMode":
        return KnxDatapointType_DPT_FanMode
    case "DPT_MasterSlaveMode":
        return KnxDatapointType_DPT_MasterSlaveMode
    case "DPT_StatusRoomSetp":
        return KnxDatapointType_DPT_StatusRoomSetp
    case "DPT_Metering_DeviceType":
        return KnxDatapointType_DPT_Metering_DeviceType
    case "DT":
        return KnxDatapointType_DT
    case "DPT_HumDehumMode":
        return KnxDatapointType_DPT_HumDehumMode
    case "DPT_EnableHCStage":
        return KnxDatapointType_DPT_EnableHCStage
    case "DPT_ADAType":
        return KnxDatapointType_DPT_ADAType
    case "DPT_BackupMode":
        return KnxDatapointType_DPT_BackupMode
    case "DPT_StartSynchronization":
        return KnxDatapointType_DPT_StartSynchronization
    case "DPT_Behaviour_Lock_Unlock":
        return KnxDatapointType_DPT_Behaviour_Lock_Unlock
    case "DPT_Behaviour_Bus_Power_Up_Down":
        return KnxDatapointType_DPT_Behaviour_Bus_Power_Up_Down
    case "DPT_DALI_Fade_Time":
        return KnxDatapointType_DPT_DALI_Fade_Time
    case "DPT_BlinkingMode":
        return KnxDatapointType_DPT_BlinkingMode
    case "DPT_LightControlMode":
        return KnxDatapointType_DPT_LightControlMode
    case "DPT_Switch":
        return KnxDatapointType_DPT_Switch
    case "DPT_SwitchPBModel":
        return KnxDatapointType_DPT_SwitchPBModel
    case "DPT_PBAction":
        return KnxDatapointType_DPT_PBAction
    case "DPT_DimmPBModel":
        return KnxDatapointType_DPT_DimmPBModel
    case "DPT_SwitchOnMode":
        return KnxDatapointType_DPT_SwitchOnMode
    case "DPT_LoadTypeSet":
        return KnxDatapointType_DPT_LoadTypeSet
    case "DPT_LoadTypeDetected":
        return KnxDatapointType_DPT_LoadTypeDetected
    case "DPT_Converter_Test_Control":
        return KnxDatapointType_DPT_Converter_Test_Control
    case "DPT_SABExcept_Behaviour":
        return KnxDatapointType_DPT_SABExcept_Behaviour
    case "DPT_SABBehaviour_Lock_Unlock":
        return KnxDatapointType_DPT_SABBehaviour_Lock_Unlock
    case "DPT_SSSBMode":
        return KnxDatapointType_DPT_SSSBMode
    case "DPT_Bool":
        return KnxDatapointType_DPT_Bool
    case "DPT_BlindsControlMode":
        return KnxDatapointType_DPT_BlindsControlMode
    case "DPT_CommMode":
        return KnxDatapointType_DPT_CommMode
    case "DPT_AddInfoTypes":
        return KnxDatapointType_DPT_AddInfoTypes
    case "DPT_RF_ModeSelect":
        return KnxDatapointType_DPT_RF_ModeSelect
    case "DPT_RF_FilterSelect":
        return KnxDatapointType_DPT_RF_FilterSelect
    case "DPT_StatusGen":
        return KnxDatapointType_DPT_StatusGen
    case "DPT_Device_Control":
        return KnxDatapointType_DPT_Device_Control
    case "DPT_ForceSign":
        return KnxDatapointType_DPT_ForceSign
    case "DPT_ForceSignCool":
        return KnxDatapointType_DPT_ForceSignCool
    case "DPT_StatusRHC":
        return KnxDatapointType_DPT_StatusRHC
    case "DPT_Enable":
        return KnxDatapointType_DPT_Enable
    case "DPT_StatusSDHWC":
        return KnxDatapointType_DPT_StatusSDHWC
    case "DPT_FuelTypeSet":
        return KnxDatapointType_DPT_FuelTypeSet
    case "DPT_StatusRCC":
        return KnxDatapointType_DPT_StatusRCC
    case "DPT_StatusAHU":
        return KnxDatapointType_DPT_StatusAHU
    case "DPT_CombinedStatus_RTSM":
        return KnxDatapointType_DPT_CombinedStatus_RTSM
    case "DPT_LightActuatorErrorInfo":
        return KnxDatapointType_DPT_LightActuatorErrorInfo
    case "DPT_RF_ModeInfo":
        return KnxDatapointType_DPT_RF_ModeInfo
    case "DPT_RF_FilterInfo":
        return KnxDatapointType_DPT_RF_FilterInfo
    case "DPT_Channel_Activation_8":
        return KnxDatapointType_DPT_Channel_Activation_8
    case "DPT_StatusDHWC":
        return KnxDatapointType_DPT_StatusDHWC
    case "WORD":
        return KnxDatapointType_WORD
    case "DPT_Ramp":
        return KnxDatapointType_DPT_Ramp
    case "DPT_StatusRHCC":
        return KnxDatapointType_DPT_StatusRHCC
    case "DPT_CombinedStatus_HVA":
        return KnxDatapointType_DPT_CombinedStatus_HVA
    case "DPT_CombinedStatus_RTC":
        return KnxDatapointType_DPT_CombinedStatus_RTC
    case "DPT_Media":
        return KnxDatapointType_DPT_Media
    case "DPT_Channel_Activation_16":
        return KnxDatapointType_DPT_Channel_Activation_16
    case "DPT_OnOffAction":
        return KnxDatapointType_DPT_OnOffAction
    case "DPT_Alarm_Reaction":
        return KnxDatapointType_DPT_Alarm_Reaction
    case "DPT_UpDown_Action":
        return KnxDatapointType_DPT_UpDown_Action
    case "DPT_HVAC_PB_Action":
        return KnxDatapointType_DPT_HVAC_PB_Action
    case "DPT_DoubleNibble":
        return KnxDatapointType_DPT_DoubleNibble
    case "DPT_Alarm":
        return KnxDatapointType_DPT_Alarm
    case "DPT_SceneInfo":
        return KnxDatapointType_DPT_SceneInfo
    case "DPT_CombinedInfoOnOff":
        return KnxDatapointType_DPT_CombinedInfoOnOff
    case "DPT_ActiveEnergy_V64":
        return KnxDatapointType_DPT_ActiveEnergy_V64
    case "DPT_ApparantEnergy_V64":
        return KnxDatapointType_DPT_ApparantEnergy_V64
    case "DPT_ReactiveEnergy_V64":
        return KnxDatapointType_DPT_ReactiveEnergy_V64
    case "DPT_Channel_Activation_24":
        return KnxDatapointType_DPT_Channel_Activation_24
    case "DPT_HVACModeNext":
        return KnxDatapointType_DPT_HVACModeNext
    case "DPT_DHWModeNext":
        return KnxDatapointType_DPT_DHWModeNext
    case "DPT_OccModeNext":
        return KnxDatapointType_DPT_OccModeNext
    case "DPT_BuildingModeNext":
        return KnxDatapointType_DPT_BuildingModeNext
    case "DPT_BinaryValue":
        return KnxDatapointType_DPT_BinaryValue
    case "DPT_Version":
        return KnxDatapointType_DPT_Version
    case "DPT_AlarmInfo":
        return KnxDatapointType_DPT_AlarmInfo
    case "DPT_TempRoomSetpSetF16_3":
        return KnxDatapointType_DPT_TempRoomSetpSetF16_3
    case "DPT_TempRoomSetpSetShiftF16_3":
        return KnxDatapointType_DPT_TempRoomSetpSetShiftF16_3
    case "DPT_Scaling_Speed":
        return KnxDatapointType_DPT_Scaling_Speed
    case "DPT_Scaling_Step_Time":
        return KnxDatapointType_DPT_Scaling_Step_Time
    case "DPT_MeteringValue":
        return KnxDatapointType_DPT_MeteringValue
    case "DPT_MBus_Address":
        return KnxDatapointType_DPT_MBus_Address
    case "DPT_Colour_RGB":
        return KnxDatapointType_DPT_Colour_RGB
    case "DPT_LanguageCodeAlpha2_ASCII":
        return KnxDatapointType_DPT_LanguageCodeAlpha2_ASCII
    case "DPT_Step":
        return KnxDatapointType_DPT_Step
    case "DPT_Tariff_ActiveEnergy":
        return KnxDatapointType_DPT_Tariff_ActiveEnergy
    case "DPT_Prioritised_Mode_Control":
        return KnxDatapointType_DPT_Prioritised_Mode_Control
    case "DPT_DALI_Control_Gear_Diagnostic":
        return KnxDatapointType_DPT_DALI_Control_Gear_Diagnostic
    case "DPT_DALI_Diagnostics":
        return KnxDatapointType_DPT_DALI_Diagnostics
    case "DPT_CombinedPosition":
        return KnxDatapointType_DPT_CombinedPosition
    case "DPT_StatusSAB":
        return KnxDatapointType_DPT_StatusSAB
    case "DPT_Colour_xyY":
        return KnxDatapointType_DPT_Colour_xyY
    case "DPT_Converter_Status":
        return KnxDatapointType_DPT_Converter_Status
    case "DPT_Converter_Test_Result":
        return KnxDatapointType_DPT_Converter_Test_Result
    case "DPT_Battery_Info":
        return KnxDatapointType_DPT_Battery_Info
    case "DPT_UpDown":
        return KnxDatapointType_DPT_UpDown
    case "DPT_Brightness_Colour_Temperature_Transition":
        return KnxDatapointType_DPT_Brightness_Colour_Temperature_Transition
    case "DPT_Brightness_Colour_Temperature_Control":
        return KnxDatapointType_DPT_Brightness_Colour_Temperature_Control
    case "DPT_Colour_RGBW":
        return KnxDatapointType_DPT_Colour_RGBW
    case "DPT_Relative_Control_RGBW":
        return KnxDatapointType_DPT_Relative_Control_RGBW
    case "DPT_Relative_Control_RGB":
        return KnxDatapointType_DPT_Relative_Control_RGB
    case "DPT_GeographicalLocation":
        return KnxDatapointType_DPT_GeographicalLocation
    case "DPT_TempRoomSetpSetF16_4":
        return KnxDatapointType_DPT_TempRoomSetpSetF16_4
    case "DPT_TempRoomSetpSetShiftF16_4":
        return KnxDatapointType_DPT_TempRoomSetpSetShiftF16_4
    case "DPT_OpenClose":
        return KnxDatapointType_DPT_OpenClose
    case "DPT_Start":
        return KnxDatapointType_DPT_Start
    case "DPT_State":
        return KnxDatapointType_DPT_State
    case "DPT_Invert":
        return KnxDatapointType_DPT_Invert
    case "DPT_DimSendStyle":
        return KnxDatapointType_DPT_DimSendStyle
    case "DWORD":
        return KnxDatapointType_DWORD
    case "DPT_InputSource":
        return KnxDatapointType_DPT_InputSource
    case "DPT_Reset":
        return KnxDatapointType_DPT_Reset
    case "DPT_Ack":
        return KnxDatapointType_DPT_Ack
    case "DPT_Trigger":
        return KnxDatapointType_DPT_Trigger
    case "DPT_Occupancy":
        return KnxDatapointType_DPT_Occupancy
    case "DPT_Window_Door":
        return KnxDatapointType_DPT_Window_Door
    case "DPT_LogicalFunction":
        return KnxDatapointType_DPT_LogicalFunction
    case "DPT_Scene_AB":
        return KnxDatapointType_DPT_Scene_AB
    case "DPT_ShutterBlinds_Mode":
        return KnxDatapointType_DPT_ShutterBlinds_Mode
    case "DPT_DayNight":
        return KnxDatapointType_DPT_DayNight
    case "LWORD":
        return KnxDatapointType_LWORD
    case "DPT_Heat_Cool":
        return KnxDatapointType_DPT_Heat_Cool
    case "DPT_Switch_Control":
        return KnxDatapointType_DPT_Switch_Control
    case "DPT_Bool_Control":
        return KnxDatapointType_DPT_Bool_Control
    case "DPT_Enable_Control":
        return KnxDatapointType_DPT_Enable_Control
    case "DPT_Ramp_Control":
        return KnxDatapointType_DPT_Ramp_Control
    case "DPT_Alarm_Control":
        return KnxDatapointType_DPT_Alarm_Control
    case "DPT_BinaryValue_Control":
        return KnxDatapointType_DPT_BinaryValue_Control
    case "DPT_Step_Control":
        return KnxDatapointType_DPT_Step_Control
    case "DPT_Direction1_Control":
        return KnxDatapointType_DPT_Direction1_Control
    case "DPT_Direction2_Control":
        return KnxDatapointType_DPT_Direction2_Control
    case "USINT":
        return KnxDatapointType_USINT
    case "DPT_Start_Control":
        return KnxDatapointType_DPT_Start_Control
    case "DPT_State_Control":
        return KnxDatapointType_DPT_State_Control
    case "DPT_Invert_Control":
        return KnxDatapointType_DPT_Invert_Control
    case "DPT_Control_Dimming":
        return KnxDatapointType_DPT_Control_Dimming
    case "DPT_Control_Blinds":
        return KnxDatapointType_DPT_Control_Blinds
    case "DPT_Char_ASCII":
        return KnxDatapointType_DPT_Char_ASCII
    case "DPT_Char_8859_1":
        return KnxDatapointType_DPT_Char_8859_1
    case "DPT_Scaling":
        return KnxDatapointType_DPT_Scaling
    case "DPT_Angle":
        return KnxDatapointType_DPT_Angle
    case "DPT_Percent_U8":
        return KnxDatapointType_DPT_Percent_U8
    case "SINT":
        return KnxDatapointType_SINT
    case "DPT_DecimalFactor":
        return KnxDatapointType_DPT_DecimalFactor
    case "DPT_Tariff":
        return KnxDatapointType_DPT_Tariff
    case "DPT_Value_1_Ucount":
        return KnxDatapointType_DPT_Value_1_Ucount
    case "DPT_FanStage":
        return KnxDatapointType_DPT_FanStage
    case "DPT_Percent_V8":
        return KnxDatapointType_DPT_Percent_V8
    case "DPT_Value_1_Count":
        return KnxDatapointType_DPT_Value_1_Count
    case "DPT_Status_Mode3":
        return KnxDatapointType_DPT_Status_Mode3
    case "DPT_Value_2_Ucount":
        return KnxDatapointType_DPT_Value_2_Ucount
    case "DPT_TimePeriodMsec":
        return KnxDatapointType_DPT_TimePeriodMsec
    case "DPT_TimePeriod10Msec":
        return KnxDatapointType_DPT_TimePeriod10Msec
    case "UINT":
        return KnxDatapointType_UINT
    case "DPT_TimePeriod100Msec":
        return KnxDatapointType_DPT_TimePeriod100Msec
    case "DPT_TimePeriodSec":
        return KnxDatapointType_DPT_TimePeriodSec
    case "DPT_TimePeriodMin":
        return KnxDatapointType_DPT_TimePeriodMin
    case "DPT_TimePeriodHrs":
        return KnxDatapointType_DPT_TimePeriodHrs
    case "DPT_PropDataType":
        return KnxDatapointType_DPT_PropDataType
    case "DPT_Length_mm":
        return KnxDatapointType_DPT_Length_mm
    case "DPT_UElCurrentmA":
        return KnxDatapointType_DPT_UElCurrentmA
    case "DPT_Brightness":
        return KnxDatapointType_DPT_Brightness
    case "DPT_Absolute_Colour_Temperature":
        return KnxDatapointType_DPT_Absolute_Colour_Temperature
    case "DPT_Value_2_Count":
        return KnxDatapointType_DPT_Value_2_Count
    case "INT":
        return KnxDatapointType_INT
    case "DPT_DeltaTimeMsec":
        return KnxDatapointType_DPT_DeltaTimeMsec
    case "DPT_DeltaTime10Msec":
        return KnxDatapointType_DPT_DeltaTime10Msec
    case "DPT_DeltaTime100Msec":
        return KnxDatapointType_DPT_DeltaTime100Msec
    case "DPT_DeltaTimeSec":
        return KnxDatapointType_DPT_DeltaTimeSec
    case "DPT_DeltaTimeMin":
        return KnxDatapointType_DPT_DeltaTimeMin
    case "DPT_DeltaTimeHrs":
        return KnxDatapointType_DPT_DeltaTimeHrs
    case "DPT_Percent_V16":
        return KnxDatapointType_DPT_Percent_V16
    case "DPT_Rotation_Angle":
        return KnxDatapointType_DPT_Rotation_Angle
    case "DPT_Length_m":
        return KnxDatapointType_DPT_Length_m
    case "DPT_Value_Temp":
        return KnxDatapointType_DPT_Value_Temp
    }
    return 0
}

func CastKnxDatapointType(structType interface{}) KnxDatapointType {
    castFunc := func(typ interface{}) KnxDatapointType {
        if sKnxDatapointType, ok := typ.(KnxDatapointType); ok {
            return sKnxDatapointType
        }
        return 0
    }
    return castFunc(structType)
}

func (m KnxDatapointType) LengthInBits() uint16 {
    return 32
}

func (m KnxDatapointType) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KnxDatapointTypeParse(io *utils.ReadBuffer) (KnxDatapointType, error) {
    val, err := io.ReadUint32(32)
    if err != nil {
        return 0, nil
    }
    return KnxDatapointTypeByValue(val), nil
}

func (e KnxDatapointType) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint32(32, uint32(e))
    return err
}

func (e KnxDatapointType) String() string {
    switch e {
    case KnxDatapointType_DPT_UNKNOWN:
        return "DPT_UNKNOWN"
    case KnxDatapointType_BOOL:
        return "BOOL"
    case KnxDatapointType_UDINT:
        return "UDINT"
    case KnxDatapointType_DPT_Value_Tempd:
        return "DPT_Value_Tempd"
    case KnxDatapointType_DPT_Value_Tempa:
        return "DPT_Value_Tempa"
    case KnxDatapointType_DPT_Value_Lux:
        return "DPT_Value_Lux"
    case KnxDatapointType_DPT_Value_Wsp:
        return "DPT_Value_Wsp"
    case KnxDatapointType_DPT_Value_Pres:
        return "DPT_Value_Pres"
    case KnxDatapointType_DPT_Value_Humidity:
        return "DPT_Value_Humidity"
    case KnxDatapointType_DPT_Value_AirQuality:
        return "DPT_Value_AirQuality"
    case KnxDatapointType_DPT_Value_AirFlow:
        return "DPT_Value_AirFlow"
    case KnxDatapointType_DPT_Value_Time1:
        return "DPT_Value_Time1"
    case KnxDatapointType_DPT_Value_Time2:
        return "DPT_Value_Time2"
    case KnxDatapointType_DINT:
        return "DINT"
    case KnxDatapointType_DPT_Value_Volt:
        return "DPT_Value_Volt"
    case KnxDatapointType_DPT_Value_Curr:
        return "DPT_Value_Curr"
    case KnxDatapointType_DPT_PowerDensity:
        return "DPT_PowerDensity"
    case KnxDatapointType_DPT_KelvinPerPercent:
        return "DPT_KelvinPerPercent"
    case KnxDatapointType_DPT_Power:
        return "DPT_Power"
    case KnxDatapointType_DPT_Value_Volume_Flow:
        return "DPT_Value_Volume_Flow"
    case KnxDatapointType_DPT_Rain_Amount:
        return "DPT_Rain_Amount"
    case KnxDatapointType_DPT_Value_Temp_F:
        return "DPT_Value_Temp_F"
    case KnxDatapointType_DPT_Value_Wsp_kmh:
        return "DPT_Value_Wsp_kmh"
    case KnxDatapointType_DPT_Value_Absolute_Humidity:
        return "DPT_Value_Absolute_Humidity"
    case KnxDatapointType_ULINT:
        return "ULINT"
    case KnxDatapointType_DPT_Concentration_ygm3:
        return "DPT_Concentration_ygm3"
    case KnxDatapointType_DPT_TimeOfDay:
        return "DPT_TimeOfDay"
    case KnxDatapointType_DPT_Date:
        return "DPT_Date"
    case KnxDatapointType_DPT_Value_4_Ucount:
        return "DPT_Value_4_Ucount"
    case KnxDatapointType_DPT_LongTimePeriod_Sec:
        return "DPT_LongTimePeriod_Sec"
    case KnxDatapointType_DPT_LongTimePeriod_Min:
        return "DPT_LongTimePeriod_Min"
    case KnxDatapointType_DPT_LongTimePeriod_Hrs:
        return "DPT_LongTimePeriod_Hrs"
    case KnxDatapointType_DPT_VolumeLiquid_Litre:
        return "DPT_VolumeLiquid_Litre"
    case KnxDatapointType_DPT_Volume_m_3:
        return "DPT_Volume_m_3"
    case KnxDatapointType_DPT_Value_4_Count:
        return "DPT_Value_4_Count"
    case KnxDatapointType_LINT:
        return "LINT"
    case KnxDatapointType_DPT_FlowRate_m3h:
        return "DPT_FlowRate_m3h"
    case KnxDatapointType_DPT_ActiveEnergy:
        return "DPT_ActiveEnergy"
    case KnxDatapointType_DPT_ApparantEnergy:
        return "DPT_ApparantEnergy"
    case KnxDatapointType_DPT_ReactiveEnergy:
        return "DPT_ReactiveEnergy"
    case KnxDatapointType_DPT_ActiveEnergy_kWh:
        return "DPT_ActiveEnergy_kWh"
    case KnxDatapointType_DPT_ApparantEnergy_kVAh:
        return "DPT_ApparantEnergy_kVAh"
    case KnxDatapointType_DPT_ReactiveEnergy_kVARh:
        return "DPT_ReactiveEnergy_kVARh"
    case KnxDatapointType_DPT_ActiveEnergy_MWh:
        return "DPT_ActiveEnergy_MWh"
    case KnxDatapointType_DPT_LongDeltaTimeSec:
        return "DPT_LongDeltaTimeSec"
    case KnxDatapointType_DPT_DeltaVolumeLiquid_Litre:
        return "DPT_DeltaVolumeLiquid_Litre"
    case KnxDatapointType_REAL:
        return "REAL"
    case KnxDatapointType_DPT_DeltaVolume_m_3:
        return "DPT_DeltaVolume_m_3"
    case KnxDatapointType_DPT_Value_Acceleration:
        return "DPT_Value_Acceleration"
    case KnxDatapointType_DPT_Value_Acceleration_Angular:
        return "DPT_Value_Acceleration_Angular"
    case KnxDatapointType_DPT_Value_Activation_Energy:
        return "DPT_Value_Activation_Energy"
    case KnxDatapointType_DPT_Value_Activity:
        return "DPT_Value_Activity"
    case KnxDatapointType_DPT_Value_Mol:
        return "DPT_Value_Mol"
    case KnxDatapointType_DPT_Value_Amplitude:
        return "DPT_Value_Amplitude"
    case KnxDatapointType_DPT_Value_AngleRad:
        return "DPT_Value_AngleRad"
    case KnxDatapointType_DPT_Value_AngleDeg:
        return "DPT_Value_AngleDeg"
    case KnxDatapointType_DPT_Value_Angular_Momentum:
        return "DPT_Value_Angular_Momentum"
    case KnxDatapointType_LREAL:
        return "LREAL"
    case KnxDatapointType_DPT_Value_Angular_Velocity:
        return "DPT_Value_Angular_Velocity"
    case KnxDatapointType_DPT_Value_Area:
        return "DPT_Value_Area"
    case KnxDatapointType_DPT_Value_Capacitance:
        return "DPT_Value_Capacitance"
    case KnxDatapointType_DPT_Value_Charge_DensitySurface:
        return "DPT_Value_Charge_DensitySurface"
    case KnxDatapointType_DPT_Value_Charge_DensityVolume:
        return "DPT_Value_Charge_DensityVolume"
    case KnxDatapointType_DPT_Value_Compressibility:
        return "DPT_Value_Compressibility"
    case KnxDatapointType_DPT_Value_Conductance:
        return "DPT_Value_Conductance"
    case KnxDatapointType_DPT_Value_Electrical_Conductivity:
        return "DPT_Value_Electrical_Conductivity"
    case KnxDatapointType_DPT_Value_Density:
        return "DPT_Value_Density"
    case KnxDatapointType_DPT_Value_Electric_Charge:
        return "DPT_Value_Electric_Charge"
    case KnxDatapointType_CHAR:
        return "CHAR"
    case KnxDatapointType_DPT_Value_Electric_Current:
        return "DPT_Value_Electric_Current"
    case KnxDatapointType_DPT_Value_Electric_CurrentDensity:
        return "DPT_Value_Electric_CurrentDensity"
    case KnxDatapointType_DPT_Value_Electric_DipoleMoment:
        return "DPT_Value_Electric_DipoleMoment"
    case KnxDatapointType_DPT_Value_Electric_Displacement:
        return "DPT_Value_Electric_Displacement"
    case KnxDatapointType_DPT_Value_Electric_FieldStrength:
        return "DPT_Value_Electric_FieldStrength"
    case KnxDatapointType_DPT_Value_Electric_Flux:
        return "DPT_Value_Electric_Flux"
    case KnxDatapointType_DPT_Value_Electric_FluxDensity:
        return "DPT_Value_Electric_FluxDensity"
    case KnxDatapointType_DPT_Value_Electric_Polarization:
        return "DPT_Value_Electric_Polarization"
    case KnxDatapointType_DPT_Value_Electric_Potential:
        return "DPT_Value_Electric_Potential"
    case KnxDatapointType_DPT_Value_Electric_PotentialDifference:
        return "DPT_Value_Electric_PotentialDifference"
    case KnxDatapointType_WCHAR:
        return "WCHAR"
    case KnxDatapointType_DPT_Value_ElectromagneticMoment:
        return "DPT_Value_ElectromagneticMoment"
    case KnxDatapointType_DPT_Value_Electromotive_Force:
        return "DPT_Value_Electromotive_Force"
    case KnxDatapointType_DPT_Value_Energy:
        return "DPT_Value_Energy"
    case KnxDatapointType_DPT_Value_Force:
        return "DPT_Value_Force"
    case KnxDatapointType_DPT_Value_Frequency:
        return "DPT_Value_Frequency"
    case KnxDatapointType_DPT_Value_Angular_Frequency:
        return "DPT_Value_Angular_Frequency"
    case KnxDatapointType_DPT_Value_Heat_Capacity:
        return "DPT_Value_Heat_Capacity"
    case KnxDatapointType_DPT_Value_Heat_FlowRate:
        return "DPT_Value_Heat_FlowRate"
    case KnxDatapointType_DPT_Value_Heat_Quantity:
        return "DPT_Value_Heat_Quantity"
    case KnxDatapointType_DPT_Value_Impedance:
        return "DPT_Value_Impedance"
    case KnxDatapointType_STRING:
        return "STRING"
    case KnxDatapointType_DPT_Value_Length:
        return "DPT_Value_Length"
    case KnxDatapointType_DPT_Value_Light_Quantity:
        return "DPT_Value_Light_Quantity"
    case KnxDatapointType_DPT_Value_Luminance:
        return "DPT_Value_Luminance"
    case KnxDatapointType_DPT_Value_Luminous_Flux:
        return "DPT_Value_Luminous_Flux"
    case KnxDatapointType_DPT_Value_Luminous_Intensity:
        return "DPT_Value_Luminous_Intensity"
    case KnxDatapointType_DPT_Value_Magnetic_FieldStrength:
        return "DPT_Value_Magnetic_FieldStrength"
    case KnxDatapointType_DPT_Value_Magnetic_Flux:
        return "DPT_Value_Magnetic_Flux"
    case KnxDatapointType_DPT_Value_Magnetic_FluxDensity:
        return "DPT_Value_Magnetic_FluxDensity"
    case KnxDatapointType_DPT_Value_Magnetic_Moment:
        return "DPT_Value_Magnetic_Moment"
    case KnxDatapointType_DPT_Value_Magnetic_Polarization:
        return "DPT_Value_Magnetic_Polarization"
    case KnxDatapointType_WSTRING:
        return "WSTRING"
    case KnxDatapointType_DPT_Value_Magnetization:
        return "DPT_Value_Magnetization"
    case KnxDatapointType_DPT_Value_MagnetomotiveForce:
        return "DPT_Value_MagnetomotiveForce"
    case KnxDatapointType_DPT_Value_Mass:
        return "DPT_Value_Mass"
    case KnxDatapointType_DPT_Value_MassFlux:
        return "DPT_Value_MassFlux"
    case KnxDatapointType_DPT_Value_Momentum:
        return "DPT_Value_Momentum"
    case KnxDatapointType_DPT_Value_Phase_AngleRad:
        return "DPT_Value_Phase_AngleRad"
    case KnxDatapointType_DPT_Value_Phase_AngleDeg:
        return "DPT_Value_Phase_AngleDeg"
    case KnxDatapointType_DPT_Value_Power:
        return "DPT_Value_Power"
    case KnxDatapointType_DPT_Value_Power_Factor:
        return "DPT_Value_Power_Factor"
    case KnxDatapointType_DPT_Value_Pressure:
        return "DPT_Value_Pressure"
    case KnxDatapointType_BYTE:
        return "BYTE"
    case KnxDatapointType_TIME:
        return "TIME"
    case KnxDatapointType_DPT_Value_Reactance:
        return "DPT_Value_Reactance"
    case KnxDatapointType_DPT_Value_Resistance:
        return "DPT_Value_Resistance"
    case KnxDatapointType_DPT_Value_Resistivity:
        return "DPT_Value_Resistivity"
    case KnxDatapointType_DPT_Value_SelfInductance:
        return "DPT_Value_SelfInductance"
    case KnxDatapointType_DPT_Value_SolidAngle:
        return "DPT_Value_SolidAngle"
    case KnxDatapointType_DPT_Value_Sound_Intensity:
        return "DPT_Value_Sound_Intensity"
    case KnxDatapointType_DPT_Value_Speed:
        return "DPT_Value_Speed"
    case KnxDatapointType_DPT_Value_Stress:
        return "DPT_Value_Stress"
    case KnxDatapointType_DPT_Value_Surface_Tension:
        return "DPT_Value_Surface_Tension"
    case KnxDatapointType_DPT_Value_Common_Temperature:
        return "DPT_Value_Common_Temperature"
    case KnxDatapointType_LTIME:
        return "LTIME"
    case KnxDatapointType_DPT_Value_Absolute_Temperature:
        return "DPT_Value_Absolute_Temperature"
    case KnxDatapointType_DPT_Value_TemperatureDifference:
        return "DPT_Value_TemperatureDifference"
    case KnxDatapointType_DPT_Value_Thermal_Capacity:
        return "DPT_Value_Thermal_Capacity"
    case KnxDatapointType_DPT_Value_Thermal_Conductivity:
        return "DPT_Value_Thermal_Conductivity"
    case KnxDatapointType_DPT_Value_ThermoelectricPower:
        return "DPT_Value_ThermoelectricPower"
    case KnxDatapointType_DPT_Value_Time:
        return "DPT_Value_Time"
    case KnxDatapointType_DPT_Value_Torque:
        return "DPT_Value_Torque"
    case KnxDatapointType_DPT_Value_Volume:
        return "DPT_Value_Volume"
    case KnxDatapointType_DPT_Value_Volume_Flux:
        return "DPT_Value_Volume_Flux"
    case KnxDatapointType_DPT_Value_Weight:
        return "DPT_Value_Weight"
    case KnxDatapointType_DATE:
        return "DATE"
    case KnxDatapointType_DPT_Value_Work:
        return "DPT_Value_Work"
    case KnxDatapointType_DPT_Volume_Flux_Meter:
        return "DPT_Volume_Flux_Meter"
    case KnxDatapointType_DPT_Volume_Flux_ls:
        return "DPT_Volume_Flux_ls"
    case KnxDatapointType_DPT_Access_Data:
        return "DPT_Access_Data"
    case KnxDatapointType_DPT_String_ASCII:
        return "DPT_String_ASCII"
    case KnxDatapointType_DPT_String_8859_1:
        return "DPT_String_8859_1"
    case KnxDatapointType_DPT_SceneNumber:
        return "DPT_SceneNumber"
    case KnxDatapointType_DPT_SceneControl:
        return "DPT_SceneControl"
    case KnxDatapointType_DPT_DateTime:
        return "DPT_DateTime"
    case KnxDatapointType_DPT_SCLOMode:
        return "DPT_SCLOMode"
    case KnxDatapointType_TIME_OF_DAY:
        return "TIME_OF_DAY"
    case KnxDatapointType_DPT_BuildingMode:
        return "DPT_BuildingMode"
    case KnxDatapointType_DPT_OccMode:
        return "DPT_OccMode"
    case KnxDatapointType_DPT_Priority:
        return "DPT_Priority"
    case KnxDatapointType_DPT_LightApplicationMode:
        return "DPT_LightApplicationMode"
    case KnxDatapointType_DPT_ApplicationArea:
        return "DPT_ApplicationArea"
    case KnxDatapointType_DPT_AlarmClassType:
        return "DPT_AlarmClassType"
    case KnxDatapointType_DPT_PSUMode:
        return "DPT_PSUMode"
    case KnxDatapointType_DPT_ErrorClass_System:
        return "DPT_ErrorClass_System"
    case KnxDatapointType_DPT_ErrorClass_HVAC:
        return "DPT_ErrorClass_HVAC"
    case KnxDatapointType_DPT_Time_Delay:
        return "DPT_Time_Delay"
    case KnxDatapointType_TOD:
        return "TOD"
    case KnxDatapointType_DPT_Beaufort_Wind_Force_Scale:
        return "DPT_Beaufort_Wind_Force_Scale"
    case KnxDatapointType_DPT_SensorSelect:
        return "DPT_SensorSelect"
    case KnxDatapointType_DPT_ActuatorConnectType:
        return "DPT_ActuatorConnectType"
    case KnxDatapointType_DPT_Cloud_Cover:
        return "DPT_Cloud_Cover"
    case KnxDatapointType_DPT_PowerReturnMode:
        return "DPT_PowerReturnMode"
    case KnxDatapointType_DPT_FuelType:
        return "DPT_FuelType"
    case KnxDatapointType_DPT_BurnerType:
        return "DPT_BurnerType"
    case KnxDatapointType_DPT_HVACMode:
        return "DPT_HVACMode"
    case KnxDatapointType_DPT_DHWMode:
        return "DPT_DHWMode"
    case KnxDatapointType_DPT_LoadPriority:
        return "DPT_LoadPriority"
    case KnxDatapointType_DATE_AND_TIME:
        return "DATE_AND_TIME"
    case KnxDatapointType_DPT_HVACContrMode:
        return "DPT_HVACContrMode"
    case KnxDatapointType_DPT_HVACEmergMode:
        return "DPT_HVACEmergMode"
    case KnxDatapointType_DPT_ChangeoverMode:
        return "DPT_ChangeoverMode"
    case KnxDatapointType_DPT_ValveMode:
        return "DPT_ValveMode"
    case KnxDatapointType_DPT_DamperMode:
        return "DPT_DamperMode"
    case KnxDatapointType_DPT_HeaterMode:
        return "DPT_HeaterMode"
    case KnxDatapointType_DPT_FanMode:
        return "DPT_FanMode"
    case KnxDatapointType_DPT_MasterSlaveMode:
        return "DPT_MasterSlaveMode"
    case KnxDatapointType_DPT_StatusRoomSetp:
        return "DPT_StatusRoomSetp"
    case KnxDatapointType_DPT_Metering_DeviceType:
        return "DPT_Metering_DeviceType"
    case KnxDatapointType_DT:
        return "DT"
    case KnxDatapointType_DPT_HumDehumMode:
        return "DPT_HumDehumMode"
    case KnxDatapointType_DPT_EnableHCStage:
        return "DPT_EnableHCStage"
    case KnxDatapointType_DPT_ADAType:
        return "DPT_ADAType"
    case KnxDatapointType_DPT_BackupMode:
        return "DPT_BackupMode"
    case KnxDatapointType_DPT_StartSynchronization:
        return "DPT_StartSynchronization"
    case KnxDatapointType_DPT_Behaviour_Lock_Unlock:
        return "DPT_Behaviour_Lock_Unlock"
    case KnxDatapointType_DPT_Behaviour_Bus_Power_Up_Down:
        return "DPT_Behaviour_Bus_Power_Up_Down"
    case KnxDatapointType_DPT_DALI_Fade_Time:
        return "DPT_DALI_Fade_Time"
    case KnxDatapointType_DPT_BlinkingMode:
        return "DPT_BlinkingMode"
    case KnxDatapointType_DPT_LightControlMode:
        return "DPT_LightControlMode"
    case KnxDatapointType_DPT_Switch:
        return "DPT_Switch"
    case KnxDatapointType_DPT_SwitchPBModel:
        return "DPT_SwitchPBModel"
    case KnxDatapointType_DPT_PBAction:
        return "DPT_PBAction"
    case KnxDatapointType_DPT_DimmPBModel:
        return "DPT_DimmPBModel"
    case KnxDatapointType_DPT_SwitchOnMode:
        return "DPT_SwitchOnMode"
    case KnxDatapointType_DPT_LoadTypeSet:
        return "DPT_LoadTypeSet"
    case KnxDatapointType_DPT_LoadTypeDetected:
        return "DPT_LoadTypeDetected"
    case KnxDatapointType_DPT_Converter_Test_Control:
        return "DPT_Converter_Test_Control"
    case KnxDatapointType_DPT_SABExcept_Behaviour:
        return "DPT_SABExcept_Behaviour"
    case KnxDatapointType_DPT_SABBehaviour_Lock_Unlock:
        return "DPT_SABBehaviour_Lock_Unlock"
    case KnxDatapointType_DPT_SSSBMode:
        return "DPT_SSSBMode"
    case KnxDatapointType_DPT_Bool:
        return "DPT_Bool"
    case KnxDatapointType_DPT_BlindsControlMode:
        return "DPT_BlindsControlMode"
    case KnxDatapointType_DPT_CommMode:
        return "DPT_CommMode"
    case KnxDatapointType_DPT_AddInfoTypes:
        return "DPT_AddInfoTypes"
    case KnxDatapointType_DPT_RF_ModeSelect:
        return "DPT_RF_ModeSelect"
    case KnxDatapointType_DPT_RF_FilterSelect:
        return "DPT_RF_FilterSelect"
    case KnxDatapointType_DPT_StatusGen:
        return "DPT_StatusGen"
    case KnxDatapointType_DPT_Device_Control:
        return "DPT_Device_Control"
    case KnxDatapointType_DPT_ForceSign:
        return "DPT_ForceSign"
    case KnxDatapointType_DPT_ForceSignCool:
        return "DPT_ForceSignCool"
    case KnxDatapointType_DPT_StatusRHC:
        return "DPT_StatusRHC"
    case KnxDatapointType_DPT_Enable:
        return "DPT_Enable"
    case KnxDatapointType_DPT_StatusSDHWC:
        return "DPT_StatusSDHWC"
    case KnxDatapointType_DPT_FuelTypeSet:
        return "DPT_FuelTypeSet"
    case KnxDatapointType_DPT_StatusRCC:
        return "DPT_StatusRCC"
    case KnxDatapointType_DPT_StatusAHU:
        return "DPT_StatusAHU"
    case KnxDatapointType_DPT_CombinedStatus_RTSM:
        return "DPT_CombinedStatus_RTSM"
    case KnxDatapointType_DPT_LightActuatorErrorInfo:
        return "DPT_LightActuatorErrorInfo"
    case KnxDatapointType_DPT_RF_ModeInfo:
        return "DPT_RF_ModeInfo"
    case KnxDatapointType_DPT_RF_FilterInfo:
        return "DPT_RF_FilterInfo"
    case KnxDatapointType_DPT_Channel_Activation_8:
        return "DPT_Channel_Activation_8"
    case KnxDatapointType_DPT_StatusDHWC:
        return "DPT_StatusDHWC"
    case KnxDatapointType_WORD:
        return "WORD"
    case KnxDatapointType_DPT_Ramp:
        return "DPT_Ramp"
    case KnxDatapointType_DPT_StatusRHCC:
        return "DPT_StatusRHCC"
    case KnxDatapointType_DPT_CombinedStatus_HVA:
        return "DPT_CombinedStatus_HVA"
    case KnxDatapointType_DPT_CombinedStatus_RTC:
        return "DPT_CombinedStatus_RTC"
    case KnxDatapointType_DPT_Media:
        return "DPT_Media"
    case KnxDatapointType_DPT_Channel_Activation_16:
        return "DPT_Channel_Activation_16"
    case KnxDatapointType_DPT_OnOffAction:
        return "DPT_OnOffAction"
    case KnxDatapointType_DPT_Alarm_Reaction:
        return "DPT_Alarm_Reaction"
    case KnxDatapointType_DPT_UpDown_Action:
        return "DPT_UpDown_Action"
    case KnxDatapointType_DPT_HVAC_PB_Action:
        return "DPT_HVAC_PB_Action"
    case KnxDatapointType_DPT_DoubleNibble:
        return "DPT_DoubleNibble"
    case KnxDatapointType_DPT_Alarm:
        return "DPT_Alarm"
    case KnxDatapointType_DPT_SceneInfo:
        return "DPT_SceneInfo"
    case KnxDatapointType_DPT_CombinedInfoOnOff:
        return "DPT_CombinedInfoOnOff"
    case KnxDatapointType_DPT_ActiveEnergy_V64:
        return "DPT_ActiveEnergy_V64"
    case KnxDatapointType_DPT_ApparantEnergy_V64:
        return "DPT_ApparantEnergy_V64"
    case KnxDatapointType_DPT_ReactiveEnergy_V64:
        return "DPT_ReactiveEnergy_V64"
    case KnxDatapointType_DPT_Channel_Activation_24:
        return "DPT_Channel_Activation_24"
    case KnxDatapointType_DPT_HVACModeNext:
        return "DPT_HVACModeNext"
    case KnxDatapointType_DPT_DHWModeNext:
        return "DPT_DHWModeNext"
    case KnxDatapointType_DPT_OccModeNext:
        return "DPT_OccModeNext"
    case KnxDatapointType_DPT_BuildingModeNext:
        return "DPT_BuildingModeNext"
    case KnxDatapointType_DPT_BinaryValue:
        return "DPT_BinaryValue"
    case KnxDatapointType_DPT_Version:
        return "DPT_Version"
    case KnxDatapointType_DPT_AlarmInfo:
        return "DPT_AlarmInfo"
    case KnxDatapointType_DPT_TempRoomSetpSetF16_3:
        return "DPT_TempRoomSetpSetF16_3"
    case KnxDatapointType_DPT_TempRoomSetpSetShiftF16_3:
        return "DPT_TempRoomSetpSetShiftF16_3"
    case KnxDatapointType_DPT_Scaling_Speed:
        return "DPT_Scaling_Speed"
    case KnxDatapointType_DPT_Scaling_Step_Time:
        return "DPT_Scaling_Step_Time"
    case KnxDatapointType_DPT_MeteringValue:
        return "DPT_MeteringValue"
    case KnxDatapointType_DPT_MBus_Address:
        return "DPT_MBus_Address"
    case KnxDatapointType_DPT_Colour_RGB:
        return "DPT_Colour_RGB"
    case KnxDatapointType_DPT_LanguageCodeAlpha2_ASCII:
        return "DPT_LanguageCodeAlpha2_ASCII"
    case KnxDatapointType_DPT_Step:
        return "DPT_Step"
    case KnxDatapointType_DPT_Tariff_ActiveEnergy:
        return "DPT_Tariff_ActiveEnergy"
    case KnxDatapointType_DPT_Prioritised_Mode_Control:
        return "DPT_Prioritised_Mode_Control"
    case KnxDatapointType_DPT_DALI_Control_Gear_Diagnostic:
        return "DPT_DALI_Control_Gear_Diagnostic"
    case KnxDatapointType_DPT_DALI_Diagnostics:
        return "DPT_DALI_Diagnostics"
    case KnxDatapointType_DPT_CombinedPosition:
        return "DPT_CombinedPosition"
    case KnxDatapointType_DPT_StatusSAB:
        return "DPT_StatusSAB"
    case KnxDatapointType_DPT_Colour_xyY:
        return "DPT_Colour_xyY"
    case KnxDatapointType_DPT_Converter_Status:
        return "DPT_Converter_Status"
    case KnxDatapointType_DPT_Converter_Test_Result:
        return "DPT_Converter_Test_Result"
    case KnxDatapointType_DPT_Battery_Info:
        return "DPT_Battery_Info"
    case KnxDatapointType_DPT_UpDown:
        return "DPT_UpDown"
    case KnxDatapointType_DPT_Brightness_Colour_Temperature_Transition:
        return "DPT_Brightness_Colour_Temperature_Transition"
    case KnxDatapointType_DPT_Brightness_Colour_Temperature_Control:
        return "DPT_Brightness_Colour_Temperature_Control"
    case KnxDatapointType_DPT_Colour_RGBW:
        return "DPT_Colour_RGBW"
    case KnxDatapointType_DPT_Relative_Control_RGBW:
        return "DPT_Relative_Control_RGBW"
    case KnxDatapointType_DPT_Relative_Control_RGB:
        return "DPT_Relative_Control_RGB"
    case KnxDatapointType_DPT_GeographicalLocation:
        return "DPT_GeographicalLocation"
    case KnxDatapointType_DPT_TempRoomSetpSetF16_4:
        return "DPT_TempRoomSetpSetF16_4"
    case KnxDatapointType_DPT_TempRoomSetpSetShiftF16_4:
        return "DPT_TempRoomSetpSetShiftF16_4"
    case KnxDatapointType_DPT_OpenClose:
        return "DPT_OpenClose"
    case KnxDatapointType_DPT_Start:
        return "DPT_Start"
    case KnxDatapointType_DPT_State:
        return "DPT_State"
    case KnxDatapointType_DPT_Invert:
        return "DPT_Invert"
    case KnxDatapointType_DPT_DimSendStyle:
        return "DPT_DimSendStyle"
    case KnxDatapointType_DWORD:
        return "DWORD"
    case KnxDatapointType_DPT_InputSource:
        return "DPT_InputSource"
    case KnxDatapointType_DPT_Reset:
        return "DPT_Reset"
    case KnxDatapointType_DPT_Ack:
        return "DPT_Ack"
    case KnxDatapointType_DPT_Trigger:
        return "DPT_Trigger"
    case KnxDatapointType_DPT_Occupancy:
        return "DPT_Occupancy"
    case KnxDatapointType_DPT_Window_Door:
        return "DPT_Window_Door"
    case KnxDatapointType_DPT_LogicalFunction:
        return "DPT_LogicalFunction"
    case KnxDatapointType_DPT_Scene_AB:
        return "DPT_Scene_AB"
    case KnxDatapointType_DPT_ShutterBlinds_Mode:
        return "DPT_ShutterBlinds_Mode"
    case KnxDatapointType_DPT_DayNight:
        return "DPT_DayNight"
    case KnxDatapointType_LWORD:
        return "LWORD"
    case KnxDatapointType_DPT_Heat_Cool:
        return "DPT_Heat_Cool"
    case KnxDatapointType_DPT_Switch_Control:
        return "DPT_Switch_Control"
    case KnxDatapointType_DPT_Bool_Control:
        return "DPT_Bool_Control"
    case KnxDatapointType_DPT_Enable_Control:
        return "DPT_Enable_Control"
    case KnxDatapointType_DPT_Ramp_Control:
        return "DPT_Ramp_Control"
    case KnxDatapointType_DPT_Alarm_Control:
        return "DPT_Alarm_Control"
    case KnxDatapointType_DPT_BinaryValue_Control:
        return "DPT_BinaryValue_Control"
    case KnxDatapointType_DPT_Step_Control:
        return "DPT_Step_Control"
    case KnxDatapointType_DPT_Direction1_Control:
        return "DPT_Direction1_Control"
    case KnxDatapointType_DPT_Direction2_Control:
        return "DPT_Direction2_Control"
    case KnxDatapointType_USINT:
        return "USINT"
    case KnxDatapointType_DPT_Start_Control:
        return "DPT_Start_Control"
    case KnxDatapointType_DPT_State_Control:
        return "DPT_State_Control"
    case KnxDatapointType_DPT_Invert_Control:
        return "DPT_Invert_Control"
    case KnxDatapointType_DPT_Control_Dimming:
        return "DPT_Control_Dimming"
    case KnxDatapointType_DPT_Control_Blinds:
        return "DPT_Control_Blinds"
    case KnxDatapointType_DPT_Char_ASCII:
        return "DPT_Char_ASCII"
    case KnxDatapointType_DPT_Char_8859_1:
        return "DPT_Char_8859_1"
    case KnxDatapointType_DPT_Scaling:
        return "DPT_Scaling"
    case KnxDatapointType_DPT_Angle:
        return "DPT_Angle"
    case KnxDatapointType_DPT_Percent_U8:
        return "DPT_Percent_U8"
    case KnxDatapointType_SINT:
        return "SINT"
    case KnxDatapointType_DPT_DecimalFactor:
        return "DPT_DecimalFactor"
    case KnxDatapointType_DPT_Tariff:
        return "DPT_Tariff"
    case KnxDatapointType_DPT_Value_1_Ucount:
        return "DPT_Value_1_Ucount"
    case KnxDatapointType_DPT_FanStage:
        return "DPT_FanStage"
    case KnxDatapointType_DPT_Percent_V8:
        return "DPT_Percent_V8"
    case KnxDatapointType_DPT_Value_1_Count:
        return "DPT_Value_1_Count"
    case KnxDatapointType_DPT_Status_Mode3:
        return "DPT_Status_Mode3"
    case KnxDatapointType_DPT_Value_2_Ucount:
        return "DPT_Value_2_Ucount"
    case KnxDatapointType_DPT_TimePeriodMsec:
        return "DPT_TimePeriodMsec"
    case KnxDatapointType_DPT_TimePeriod10Msec:
        return "DPT_TimePeriod10Msec"
    case KnxDatapointType_UINT:
        return "UINT"
    case KnxDatapointType_DPT_TimePeriod100Msec:
        return "DPT_TimePeriod100Msec"
    case KnxDatapointType_DPT_TimePeriodSec:
        return "DPT_TimePeriodSec"
    case KnxDatapointType_DPT_TimePeriodMin:
        return "DPT_TimePeriodMin"
    case KnxDatapointType_DPT_TimePeriodHrs:
        return "DPT_TimePeriodHrs"
    case KnxDatapointType_DPT_PropDataType:
        return "DPT_PropDataType"
    case KnxDatapointType_DPT_Length_mm:
        return "DPT_Length_mm"
    case KnxDatapointType_DPT_UElCurrentmA:
        return "DPT_UElCurrentmA"
    case KnxDatapointType_DPT_Brightness:
        return "DPT_Brightness"
    case KnxDatapointType_DPT_Absolute_Colour_Temperature:
        return "DPT_Absolute_Colour_Temperature"
    case KnxDatapointType_DPT_Value_2_Count:
        return "DPT_Value_2_Count"
    case KnxDatapointType_INT:
        return "INT"
    case KnxDatapointType_DPT_DeltaTimeMsec:
        return "DPT_DeltaTimeMsec"
    case KnxDatapointType_DPT_DeltaTime10Msec:
        return "DPT_DeltaTime10Msec"
    case KnxDatapointType_DPT_DeltaTime100Msec:
        return "DPT_DeltaTime100Msec"
    case KnxDatapointType_DPT_DeltaTimeSec:
        return "DPT_DeltaTimeSec"
    case KnxDatapointType_DPT_DeltaTimeMin:
        return "DPT_DeltaTimeMin"
    case KnxDatapointType_DPT_DeltaTimeHrs:
        return "DPT_DeltaTimeHrs"
    case KnxDatapointType_DPT_Percent_V16:
        return "DPT_Percent_V16"
    case KnxDatapointType_DPT_Rotation_Angle:
        return "DPT_Rotation_Angle"
    case KnxDatapointType_DPT_Length_m:
        return "DPT_Length_m"
    case KnxDatapointType_DPT_Value_Temp:
        return "DPT_Value_Temp"
    }
    return ""
}