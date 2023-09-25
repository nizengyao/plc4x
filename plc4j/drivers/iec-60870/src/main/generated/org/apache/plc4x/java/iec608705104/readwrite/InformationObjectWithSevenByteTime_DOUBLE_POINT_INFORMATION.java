/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.iec608705104.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION
    extends InformationObjectWithSevenByteTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.DOUBLE_POINT_INFORMATION_WITH_TIME_TAG_CP56TIME2A;
  }

  // Properties.
  protected final DoublePointInformation diq;
  protected final SevenOctetBinaryTime cp56Time2a;

  public InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION(
      int address, DoublePointInformation diq, SevenOctetBinaryTime cp56Time2a) {
    super(address);
    this.diq = diq;
    this.cp56Time2a = cp56Time2a;
  }

  public DoublePointInformation getDiq() {
    return diq;
  }

  public SevenOctetBinaryTime getCp56Time2a() {
    return cp56Time2a;
  }

  @Override
  protected void serializeInformationObjectWithSevenByteTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION");

    // Simple Field (diq)
    writeSimpleField(
        "diq",
        diq,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (cp56Time2a)
    writeSimpleField(
        "cp56Time2a",
        cp56Time2a,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (diq)
    lengthInBits += diq.getLengthInBits();

    // Simple field (cp56Time2a)
    lengthInBits += cp56Time2a.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithSevenByteTimeBuilder
      staticParseInformationObjectWithSevenByteTimeBuilder(
          ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
          throws ParseException {
    readBuffer.pullContext("InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    DoublePointInformation diq =
        readSimpleField(
            "diq",
            new DataReaderComplexDefault<>(
                () -> DoublePointInformation.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    SevenOctetBinaryTime cp56Time2a =
        readSimpleField(
            "cp56Time2a",
            new DataReaderComplexDefault<>(
                () -> SevenOctetBinaryTime.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION");
    // Create the instance
    return new InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATIONBuilderImpl(
        diq, cp56Time2a);
  }

  public static class InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATIONBuilderImpl
      implements InformationObjectWithSevenByteTime.InformationObjectWithSevenByteTimeBuilder {
    private final DoublePointInformation diq;
    private final SevenOctetBinaryTime cp56Time2a;

    public InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATIONBuilderImpl(
        DoublePointInformation diq, SevenOctetBinaryTime cp56Time2a) {
      this.diq = diq;
      this.cp56Time2a = cp56Time2a;
    }

    public InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION build(int address) {
      InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION
          informationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION =
              new InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION(
                  address, diq, cp56Time2a);
      return informationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION)) {
      return false;
    }
    InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION that =
        (InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION) o;
    return (getDiq() == that.getDiq())
        && (getCp56Time2a() == that.getCp56Time2a())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDiq(), getCp56Time2a());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}